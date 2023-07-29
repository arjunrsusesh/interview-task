// Package api provides helpers for building an API server.
package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Serve(router http.Handler, port string) {
	s := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       120 * time.Second,
		WriteTimeout:      0,
		IdleTimeout:       60 * time.Second,
	}
	StartServer(s)
}

var ShutdownTimeout time.Duration = 5 * time.Second

func StartServer(s *http.Server) {
	shutdownComplete := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Println("HTTP server Shutdown: ")
			os.Exit(1)
		}
		close(shutdownComplete)
	}()
	log.Println("HTTP server listening to ")
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe:")
	}
	<-shutdownComplete
}
