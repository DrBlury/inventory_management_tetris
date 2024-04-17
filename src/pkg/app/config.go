package app

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	"linuxcode/inventory_manager/pkg/repo"
	"linuxcode/inventory_manager/pkg/server"
	"linuxcode/inventory_manager/pkg/server/router"
	"linuxcode/inventory_manager/pkg/service/cache"
	"linuxcode/inventory_manager/pkg/telemetry"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Info        *domain.Info
	Router      *router.Config
	Server      *server.Config
	Database    *repo.Config
	OTelConfig  *telemetry.Config
	CacheConfig *cache.Config
}

func SetDefaults() {
	viper.SetDefault("APP_SERVER_PORT", "80")
	viper.SetDefault("APP_SERVER_TIMEOUT", 60*time.Second)
	viper.SetDefault("APP_SERVER_CORS_HEADERS", []string{"*"})
	viper.SetDefault("APP_SERVER_CORS_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	viper.SetDefault("APP_SERVER_CORS_ORIGINS", []string{"*"})
}

func LoadConfig(
	version string,
	buildDate string,
	description string,
	commitHash string,
	commitDate string,
) (*Config, error) {
	SetDefaults()
	viper.AutomaticEnv()

	infoConfig := &domain.Info{
		Version:     version,
		BuildDate:   buildDate,
		Description: description,
		CommitHash:  commitHash,
		CommitDate:  commitDate,
	}

	routerConfig := &router.Config{
		Timeout: viper.GetDuration("APP_SERVER_TIMEOUT"),
		CORS: router.CORSConfig{
			AllowCredentials: viper.GetBool("APP_SERVER_CORS_ALLOW_CREDENTIALS"),
			Headers:          viper.GetStringSlice("APP_SERVER_CORS_HEADERS"),
			Methods:          viper.GetStringSlice("APP_SERVER_CORS_METHODS"),
			Origins:          viper.GetStringSlice("APP_SERVER_CORS_ORIGINS"),
		},
	}

	databaseConfig := &repo.Config{
		Host:         viper.GetString("APP_DATABASE_HOST"),
		Port:         viper.GetInt("APP_DATABASE_PORT"),
		DatabaseName: viper.GetString("APP_DATABASE_NAME"),
		Username:     viper.GetString("APP_DATABASE_USER"),
		Password:     viper.GetString("APP_DATABASE_PASSWORD"),
		Level:        viper.GetString("APP_DATABASE_LOG_LEVEL"),
	}

	serverConfig := &server.Config{
		Address: "0.0.0.0:" + viper.GetString("APP_SERVER_PORT"),
		BaseURL: viper.GetString("APP_SERVER_BASE_URL"),
	}

	otelConfig := &telemetry.Config{
		EnableOTel:     viper.GetBool("OTEL_ENABLED"),
		MeterProvider:  viper.GetString("OTEL_METER_PROVIDER"),
		TracerProvider: viper.GetString("OTEL_TRACER_PROVIDER"),
	}

	cacheConfig := &cache.Config{
		Host: viper.GetString("APP_CACHE_HOST"),
		Port: viper.GetInt("APP_CACHE_PORT"),
	}

	return &Config{
		Info:        infoConfig,
		Router:      routerConfig,
		Server:      serverConfig,
		Database:    databaseConfig,
		OTelConfig:  otelConfig,
		CacheConfig: cacheConfig,
	}, nil
}
