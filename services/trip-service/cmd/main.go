package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	fare := &domain.RideFareModel{
		UserID: "42",
	}

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	t, err := svc.CreateTrip(ctx, fare)

	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	for {
		time.Sleep(time.Second)
	}
}
