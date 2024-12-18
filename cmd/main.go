package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ercross/cer/api"
	"github.com/ercross/cer/config"
	exchange "github.com/ercross/cer/internal/services/exchange_rate_provider"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {

	awsCredentials := os.Getenv("AWS_CREDENTIALS")
	if awsCredentials == "" {
		return errors.New("AWS_CREDENTIALS environment variable not set")
	}
	cfg, err := config.LoadConfig(awsCredentials)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	providerA := exchange.NewProviderA(cfg.ProviderAApiKey)
	providerB := exchange.NewProviderB(cfg.ProviderBApiKey)

	srv := api.NewServer(providerA, providerB)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", cfg.ApiPort),
		Handler: srv,
	}

	// graceful shutdown
	go func() {
		log.Printf("listening on port %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("error listening and serving: %s\n", err.Error())
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down http server: %s\n", err.Error())
		}
	}()
	wg.Wait()
	return nil
}
