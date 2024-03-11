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
	// ===== Logger =====
	logger := logging.SetLogger()

	// ===== OpenTelemetry =====
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := telemetry.SetupOTelSDK(ctx)
	if err != nil {
		return err
	}
	logger.Info("OpenTelemetry SDK initialized")
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// ===== Database =====
	db, err := repo.CreateDB(cfg.Database)
	if err != nil {
		logger.Error("error connecting to database", zap.Error(err))
		return err
	}
	logger.Info("database connection established")

	// ===== App Logic =====
	appLogic := domain.NewAppLogic(db, logger)
	logger.Info("app logic initialized")

	// ===== Handlers =====
	versionInfo := apihandler.VersionInfo{
		Version:     cfg.Info.Version,
		BuildDate:   cfg.Info.BuildDate,
		Description: cfg.Info.Description,
		CommitHash:  cfg.Info.CommitHash,
		CommitDate:  cfg.Info.CommitDate,
	}

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := apihandler.NewAPIHandler(appLogic, versionInfo, logger)
	logger.Info("api handler initialized")
	// ===== Router =====
	r := router.New(apiHandler, cfg.Router)
	logger.Info("router initialized")

	// ===== Server =====
	srv := server.NewServer(cfg.Server, r)
	logger.Info("server initialized")

	srvErr := make(chan error, 1)
	go func() {
		logger.Info("server started", zap.String("address", cfg.Server.Address))
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
