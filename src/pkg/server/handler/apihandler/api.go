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
	BaseURL  string
}

func NewAPIHandler(appLogic usecase.AppLogic, info *domain.Info, logger *zap.SugaredLogger, baseURL string) *APIHandler {
	return &APIHandler{
		AppLogic: appLogic,
		Info:     info,
		log:      logger,
		BaseURL:  baseURL,
	}
}
