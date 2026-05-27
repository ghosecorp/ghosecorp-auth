package config

import "os"

type Config struct {
	EmailURL string
	APIKey   string
}

func Load() *Config {
	return &Config{
		EmailURL: os.Getenv("BREVO_URL"),
		APIKey:   os.Getenv("BREVO_API_KEY"),
	}
}
