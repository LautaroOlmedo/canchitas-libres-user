package main

import (
	"canchitas-libres-user/internal/configuration"
	database2 "canchitas-libres-user/internal/database"
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"canchitas-libres-user/internal/pkg/infrastructure/respository/storage"
	"canchitas-libres-user/internal/pkg/infrastructure/web"
	"context"
	"fmt"
)

func main() {
	config, err := configuration.Load("../.env")
	if err != nil {
		panic(err)
	}

	//database connection
	database, err := database2.NewDBConnection(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// repository layer
	//sliceRepository := storage.NewSliceStorage(config)
	postgresStorage := storage.NewPostgresStorage(config, database)

	// application layer (services layer)
	service := domain.NewService(config, postgresStorage)

	// infrastructure layer
	handler := web.NewHandler(service)
	serv, err := web.NewServer(config, handler)
	if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		panic(err)
	}
	// Start application
	serv.Start()
}
