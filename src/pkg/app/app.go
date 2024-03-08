package app

import (
	"context"
	"errors"
	"linuxcode/inventory_manager/pkg/domain"
	"linuxcode/inventory_manager/pkg/repo"
	"linuxcode/inventory_manager/pkg/server"
	"linuxcode/inventory_manager/pkg/server/handler/apihandler"
	"linuxcode/inventory_manager/pkg/server/handler/infohandler"
	"linuxcode/inventory_manager/pkg/server/router"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

// Run runs the linuxcode/inventory_managerlication
// nolint: funlen
func Run(cfg *Config, shutdownChannel chan os.Signal) error {
	// ===== Database =====
	db, err := repo.CreateDB(cfg.Database)
	if err != nil {
		return err
	}

	// ===== App Logic =====
	appLogic := domain.NewAppLogic(db)

	// ===== Handlers =====
	versionHandler := infohandler.NewVersionHandler(
		cfg.Info.Version,
		cfg.Info.BuildDate,
		cfg.Info.Description,
		cfg.Info.CommitHash,
		cfg.Info.CommitDate,
	)

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := apihandler.NewAPIHandler(appLogic)

	// ===== Router =====
	r := router.New(versionHandler, apiHandler, cfg.Router)

	// ===== Server =====
	srv := server.NewServer(cfg.Server, r)

	// let server serve with the given router
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("server error", zap.Error(err))
		}
	}()

	zap.L().Info("server started", zap.String("address", cfg.Server.Address))

	<-shutdownChannel

	zap.L().Info("shutting down server")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	return srv.Shutdown(ctxShutDown)
}
