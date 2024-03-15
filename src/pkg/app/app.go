package app

import (
	"context"
	"errors"
	"linuxcode/inventory_manager/pkg/domain"
	"linuxcode/inventory_manager/pkg/logging"
	"linuxcode/inventory_manager/pkg/repo"
	"linuxcode/inventory_manager/pkg/server"
	"linuxcode/inventory_manager/pkg/server/handler/apihandler"
	"linuxcode/inventory_manager/pkg/server/router"
	"linuxcode/inventory_manager/pkg/telemetry"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

// Run runs the linuxcode/inventory_managerlication
// nolint: funlen
func Run(cfg *Config, shutdownChannel chan os.Signal) error {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// ===== Logger =====
	logger := logging.SetLogger()

	// ===== OpenTelemetry =====
	if cfg.OTelConfig.EnableOTel {
		logger.Info("OpenTelemetry is enabled")
		otelService := telemetry.NewOtelService(cfg.OTelConfig, ctx)
		otelShutdown, err := otelService.SetupOTelSDK(ctx)
		if err != nil {
			return err
		}
		// Handle shutdown properly so nothing leaks.
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
		logger.Info("OpenTelemetry SDK initialized")
	}
	// ===== Database =====
	db, err := repo.CreateDB(cfg.Database)
	if err != nil {
		logger.Error("error connecting to database", zap.Error(err))
		return err
	}
	logger.Info("database connection established")

	// ===== App Logic =====
	appLogic := domain.NewAppLogic(db, logger)

	// ===== Handlers =====
	apiHandler := apihandler.NewAPIHandler(appLogic, cfg.Info, logger)

	// ===== Router =====
	r := router.New(apiHandler, cfg.Router)

	// ===== Server =====
	srv := server.NewServer(cfg.Server, r)

	srvErr := make(chan error, 1)
	go func() {
		logger.Info("server started, address: ", cfg.Server.Address)
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		logger.Fatal("server error", zap.Error(err))
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		err := srv.Shutdown(context.Background())
		if err != nil {
			logger.Fatal("server shutdown error", zap.Error(err))
			return err
		}
		stop()
	}
	return nil
}
