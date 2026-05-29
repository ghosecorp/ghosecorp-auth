package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EmailURL string
	APIKey   string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("the error is: ", err)
	}

	return &Config{
		EmailURL: os.Getenv("BREVO_URL"),
		APIKey:   os.Getenv("BREVO_API_KEY"),
	}
}
