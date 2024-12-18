package v1

import (
	"errors"
	"github.com/ercross/cer/api/utils"
	exchange "github.com/ercross/cer/internal/services/exchange_rate_provider"
	"log"
	"math"
	"net/http"
)

type exchangeRateProviderResult struct {
	result float64
	err    error
}

func exchangeRate(providerA exchange.RateProvider, providerB exchange.RateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pair := r.URL.Query().Get("pair")
		if err := isValidCurrencyPair(pair); err != nil {
			utils.SendApiErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		chanA := make(chan exchangeRateProviderResult)
		chanB := make(chan exchangeRateProviderResult)

		taskA := func() (float64, error) {
			return providerA.FetchExchangeRate(pair)
		}
		taskB := func() (float64, error) {
			return providerB.FetchExchangeRate(pair)
		}

		responseSent := false

		go worker(taskA, chanA)
		go worker(taskB, chanB)

		var result exchangeRateProviderResult
		select {
		case result = <-chanA:
			if result.err != nil {
				secondResult := <-chanB
				responseSent = handleResult(secondResult, pair, w, providerB.Name())
			} else {
				responseSent = handleResult(result, pair, w, providerA.Name())
			}

		case result = <-chanB:
			if result.err != nil {
				secondResult := <-chanA
				responseSent = handleResult(secondResult, pair, w, providerA.Name())
			} else {
				responseSent = handleResult(result, pair, w, providerB.Name())
			}
		}

		if !responseSent {
			utils.SendApiErrorResponse(w, http.StatusServiceUnavailable, errors.New("rate provider not available"))
		}
	}
}

func worker(task func() (float64, error), ch chan<- exchangeRateProviderResult) {
	defer close(ch)
	result, err := task()
	ch <- exchangeRateProviderResult{result: result, err: err}
}

func isValidCurrencyPair(pair string) error {
	// do some validation for each currency in the pair
	return nil
}

// handleResult and return a boolean to indicate if response was sent
func handleResult(result exchangeRateProviderResult, pair string, w http.ResponseWriter, providerName string) bool {
	if result.err != nil {
		log.Printf("%s: %v", providerName, result.err)
		return false
	} else {
		utils.SendApiResponse(w, 200, map[string]interface{}{
			pair: convertToTwoDecimal(result.result),
		})
		return true
	}
}

func convertToTwoDecimal(value float64) float64 {
	return math.Round(value*100) / 100
}
