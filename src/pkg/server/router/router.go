package router

import (
	"fmt"
	server "linuxcode/inventory_manager/pkg/server/generated"
	apihandler "linuxcode/inventory_manager/pkg/server/handler/api"
	"linuxcode/inventory_manager/pkg/server/handler/infohandler"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	oapiMW "github.com/oapi-codegen/nethttp-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// New returns a new *chi.Mux to be used for http request routing.
func New(
	versionHandle *infohandler.VersionHandler,
	apiHandle *apihandler.APIHandler,
	cfg *Config,
) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.Timeout))
	
	// Allow CORS globally
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CORS.Origins, // Use this to allow specific origin hosts
		AllowedMethods:   cfg.CORS.Methods,
		AllowedHeaders:   cfg.CORS.Headers,
		AllowCredentials: cfg.CORS.AllowCredentials,
	}))
	
	router.Route("/info", func(r chi.Router) {
		r.Get("/health", infohandler.HealthCheck)
		r.Get("/version", versionHandle.VersionCheck)
		r.Handle("/metrics", promhttp.Handler())
	})
	
	// Log requests apart from /info
	handlerGroup := router.Group(nil)
	handlerGroup.Use(
		middleware.RequestLogger(
			&middleware.DefaultLogFormatter{
				Logger:  log.StandardLogger(),
				NoColor: false,
			},
		),
	)
	
	swagger, err := server.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	router.Use(oapiMW.OapiRequestValidator(swagger))
	
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := apihandler.NewAPIHandler()

	server.HandlerFromMux(apiHandler, router)

	return router
}
