package config

import "os"

type Config struct {
	Port         string
	DatabaseURL  string
	CookieSecure bool
}

func Load() Config {
	return Config{
		Port:         getenv("PORT", "8001"),
		DatabaseURL:  getenv("DATABASE_URL", ""),
		CookieSecure: getenv("APP_ENV", "dev") == "prod",
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
