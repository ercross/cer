package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of exchange.RateProvider
type mockRateProvider struct {
	rate float64
	err  error
	name string
}

func (m *mockRateProvider) FetchExchangeRate(pair string) (float64, error) {
	return m.rate, m.err
}

func (m *mockRateProvider) Name() string {
	return m.name
}

func TestExchangeRateHandler(t *testing.T) {
	tests := []struct {
		name              string
		pair              string
		providerAResponse exchangeRateProviderResult
		providerBResponse exchangeRateProviderResult
		expectedStatus    int
		expectedBody      map[string]interface{}
	}{
		{
			name: "Valid response from providerA",
			pair: "USD-NGN",
			providerAResponse: exchangeRateProviderResult{
				result: 756.34,
				err:    nil,
			},
			providerBResponse: exchangeRateProviderResult{
				result: 0,
				err:    errors.New("provider B error"),
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"USD-NGN": 756.34},
		},
		{
			name: "Valid response from providerB",
			pair: "EUR-USD",
			providerAResponse: exchangeRateProviderResult{
				result: 0,
				err:    errors.New("provider A error"),
			},
			providerBResponse: exchangeRateProviderResult{
				result: 1.13,
				err:    nil,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"EUR-USD": 1.13},
		},
		{
			name: "Both providers fail",
			pair: "GBP-USD",
			providerAResponse: exchangeRateProviderResult{
				result: 0,
				err:    errors.New("provider A error"),
			},
			providerBResponse: exchangeRateProviderResult{
				result: 0,
				err:    errors.New("provider B error"),
			},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock providers
			mockProviderA := &mockRateProvider{
				rate: tt.providerAResponse.result,
				err:  tt.providerAResponse.err,
				name: "Provider A",
			}
			mockProviderB := &mockRateProvider{
				rate: tt.providerBResponse.result,
				err:  tt.providerBResponse.err,
				name: "Provider B",
			}

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/exchange?pair="+tt.pair, nil)
			respRecorder := httptest.NewRecorder()

			// Call the handler
			handler := exchangeRate(mockProviderA, mockProviderB)
			handler.ServeHTTP(respRecorder, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, respRecorder.Code)

			// Assert response body if expectedBody is not nil
			if tt.expectedBody != nil {
				var actualBody map[string]interface{}
				err := json.Unmarshal(respRecorder.Body.Bytes(), &actualBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, actualBody)
			}
		})
	}
}
