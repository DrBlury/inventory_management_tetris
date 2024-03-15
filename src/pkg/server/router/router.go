package router

import (
	"fmt"
	server "linuxcode/inventory_manager/pkg/server/generated"
	"linuxcode/inventory_manager/pkg/server/handler/apihandler"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	oapiMW "github.com/oapi-codegen/nethttp-middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// New returns a new *chi.Mux to be used for http request routing.
func New(
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

	router.Use(oapiMW.OapiRequestValidator(swagger))
	// register the APIHandler
	// create handler from mux
	// and mount it to the router
	apiHandler := server.HandlerFromMux(apiHandle, router)
	oTelHandler := otelhttp.NewHandler(apiHandler, "/")

	// What is a mux?
	// A mux is an HTTP request multiplexer.
	// It matches the URL of each incoming request
	// against a list of registered patterns and calls \
	// the handler for the pattern that most
	// closely matches the URL.

	// register subrouter to the main router
	router.Mount("/", oTelHandler)

	return router
}
