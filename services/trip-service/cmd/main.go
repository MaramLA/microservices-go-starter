package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"

	grpcServer "google.golang.org/grpc"
)

var httpTripAddr = env.GetString("HTTP_TRIP_ADDR", ":8083")

var GrpcAddr = ":9093"

func main() {
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen %v ", err)
	}

	grpcServer := grpcServer.NewServer()

	// TODO: initialize gRPC handler implementation

	log.Printf("starting gRPC server trip service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	log.Println("shutting down the server")
	grpcServer.GracefulStop()

}
