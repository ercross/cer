package api

import (
	v1 "github.com/ercross/cer/api/v1"
	"github.com/ercross/cer/internal/services/exchange_rate_provider"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewServer(providerA exchange.RateProvider, providerB exchange.RateProvider) http.Handler {
	mux := chi.NewRouter()

	mux.Mount("/api/v1", v1.AddRoutes(providerA, providerB))

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	return mux
}
