package router

import (
	"fmt"
	server "linuxcode/inventory_manager/pkg/server/generated"
	apihandler "linuxcode/inventory_manager/pkg/server/handler/apihandler"
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
	swagger, err := server.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
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

	apiRouter := chi.NewRouter()
	apiRouter.Use(oapiMW.OapiRequestValidator(swagger))
	// register the APIHandler
	// create handler from mux
	// and mount it to the router
	apiHandler := server.HandlerFromMux(apiHandle, apiRouter)
	
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
	
	router.Route("/info", func(r chi.Router) {
		r.Get("/health", infohandler.HealthCheck)
		r.Get("/version", versionHandle.VersionCheck)
		r.Handle("/metrics", promhttp.Handler())
	})

	// register subrouter to the main router
	router.Mount("/", apiHandler)

	return router
}
