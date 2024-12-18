package v1

import (
	exchange "github.com/ercross/cer/internal/services/exchange_rate_provider"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func AddRoutes(providerA exchange.RateProvider, providerB exchange.RateProvider) http.Handler {
	router := chi.NewRouter()
	router.Mount("/exchange", addCurrencyExchangeRateRoutes(providerA, providerB))
	return router
}

func addCurrencyExchangeRateRoutes(providerA exchange.RateProvider, providerB exchange.RateProvider) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Get("/", exchangeRate(providerA, providerB))
	return router
}
