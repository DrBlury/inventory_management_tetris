package handler

import (
	"linuxcode/inventory_manager/pkg/model/dto"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// HandleInternalServerError is a convenient method to log and handle internal server errors.
func HandleInternalServerError(w http.ResponseWriter, r *http.Request, err error, logMsg ...string) {
	uniqueErrID := uuid.New().String()

	render.Status(r, http.StatusInternalServerError)
	render.Respond(w, r, dto.ErrorInternalServerError{
		Msg:  "internal server error",
		Uuid: uniqueErrID,
	})
}
