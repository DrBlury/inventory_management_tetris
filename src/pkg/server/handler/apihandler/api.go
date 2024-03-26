package apihandler

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	"linuxcode/inventory_manager/pkg/domain/usecase"

	"go.uber.org/zap"
)

type APIHandler struct {
	AppLogic usecase.AppLogic
	Info     *domain.Info
	log      *zap.SugaredLogger
}

func NewAPIHandler(appLogic usecase.AppLogic, info *domain.Info, logger *zap.SugaredLogger) *APIHandler {
	return &APIHandler{
		AppLogic: appLogic,
		Info:     info,
		log:      logger,
	}
}
