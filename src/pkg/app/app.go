package app

import (
	"context"
	"errors"
	"linuxcode/inventory_manager/pkg/domain"
	"linuxcode/inventory_manager/pkg/server"
	"linuxcode/inventory_manager/pkg/server/handler/apihandler"
	"linuxcode/inventory_manager/pkg/server/handler/infohandler"
	"linuxcode/inventory_manager/pkg/server/router"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Run runs the linuxcode/inventory_managerlication
// nolint: funlen
func Run(cfg *Config, shutdownChannel chan os.Signal) error {
	versionHandler := infohandler.NewVersionHandler(
		cfg.Info.Version,
		cfg.Info.BuildDate,
		cfg.Info.Description,
		cfg.Info.CommitHash,
		cfg.Info.CommitDate,
	)

	// create app logic layer
	appLogic := domain.NewAppLogic()

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := apihandler.NewAPIHandler(appLogic)

	// setup router that uses the handlers
	r := router.New(versionHandler, apiHandler, cfg.Router)

	// Set up server
	srv := server.NewServer(cfg.Server, r)

	// let server serve with the given router
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("http server stopped without being closed explicitly")
		}
	}()

	log.WithField("addr", cfg.Server.Address).Info("server started listening")

	<-shutdownChannel

	log.Info("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	return srv.Shutdown(ctxShutDown)
}
