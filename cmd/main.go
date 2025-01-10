package main

import (
	"canchitas-libres-field/internal/configuration"
	database2 "canchitas-libres-field/internal/database"
	"canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/infrastructure/respository/storage"
	web2 "canchitas-libres-field/internal/pkg/infrastructure/web"
	"context"
	"fmt"
)

func main() {
	config, err := configuration.Load("../.env")
	if err != nil {
		panic(err)
	}

	// database connection
	database, err := database2.NewDBConnection(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// repository layer
	//sliceRepository := storage.NewSliceStorage(config)
	pgStorage := storage.NewPostgresStorage(config, database)

	// application layer (services layer)
	service := domain.NewService(config, pgStorage) // acá ocurre inyección de dependencia

	// infrastructure layer
	handler := web2.NewHandler(service)          // acá ocurre inyección de dependencia
	serv, err := web2.NewServer(config, handler) // acá ocurre inyección de dependencia
	if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		panic(err)
	}
	// Start application
	serv.Start()
}
