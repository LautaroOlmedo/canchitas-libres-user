package main

import (
	"canchitas-libres-field/internal/configuration"
	"canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/respository/storage"
	"canchitas-libres-field/internal/pkg/web"
	"fmt"
)

func main() {
	config, err := configuration.Load("../.env")
	if err != nil {
		panic(err)
	}
	sliceRepository := storage.NewSliceStorage(config)

	// application layer (services layer)
	service := domain.NewService(config, sliceRepository)

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
