package app

import (
	"linuxcode/inventory_manager/pkg/server"
	"linuxcode/inventory_manager/pkg/server/router"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Info   *Info
	Router *router.Config
	Server *server.Config
	// Add your configs below here
}

// Info is configurable information usually set at build time with ldflags.
type Info struct {
	Version     string
	BuildDate   string
	Description string
	CommitHash  string
	CommitDate  string
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

	infoConfig := &Info{
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

	serverConfig := &server.Config{
		Address: "0.0.0.0:" + viper.GetString("APP_SERVER_PORT"),
	}

	// Add your configs below here

	return &Config{
		Info:   infoConfig,
		Router: routerConfig,
		Server: serverConfig,
		// Add your config below here
	}, nil
}
