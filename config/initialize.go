package config

import (
	"errors"
	"os"
)

func NewAppConfig(port string, databaseUrl, adminUsername, adminPassword string) *AppConfig {
	return &AppConfig{
		Port:          port,
		DatabaseUrl:   databaseUrl,
		AdminUsername: adminUsername,
		AdminPassword: adminPassword,
	}
}

func NewAppConfigFromEnv() (*AppConfig, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT not found in environment variable")
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return nil, errors.New("DATABASE_URL not found in environment variable")
	}

	adminUsername := os.Getenv("ADMIN_USERNAME")

	if adminUsername == "" {
		return nil, errors.New("ADMIN_USERNAME not found in environment variable")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminPassword == "" {
		return nil, errors.New("ADMIN_PASSWORD not found in environment variable")
	}

	return NewAppConfig(
		port,
		databaseUrl,
		adminUsername,
		adminPassword,
	), nil
}
