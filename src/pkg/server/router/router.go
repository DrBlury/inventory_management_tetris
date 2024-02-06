package router

import (
	"linuxcode/inventory_manager/pkg/server/handler/infohandler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// New returns a new *chi.Mux to be used for http request routing.
func New(
	versionHandle *infohandler.VersionHandler,
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

	handlerGroup.Route("/api", func(r chi.Router) {
		// Add your own routes here
	})

	return router
}
