package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Server struct {
	DOMAIN      string `yaml:"DOMAIN"`
	SERVER_PORT string `yaml:"SERVER_PORT"`
}

type Configuration struct {
	SERVER Server `yaml:"SERVER"`
}

func Load(file string) (*Configuration, error) {
	err := godotenv.Load(file)
	if err != nil {
		fmt.Println(".env file not found. Operating system environment variables will be loaded.")
	}
	return &Configuration{
		SERVER: Server{
			DOMAIN:      os.Getenv("SERVER_DOMAIN"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
	}, nil
}
