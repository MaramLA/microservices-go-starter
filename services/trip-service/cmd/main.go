package main

import (
	"log"
	"net/http"

	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
)

var httpTripAddr = env.GetString("HTTP_TRIP_ADDR", ":8083")

func main() {
	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	mux := http.NewServeMux()

	httpHandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpTripAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error %v", err)
	}
}

