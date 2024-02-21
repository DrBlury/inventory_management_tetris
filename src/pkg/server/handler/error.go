package handler

import (
	server "linuxcode/inventory_manager/pkg/server/generated"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// HandleInternalServerError is a convenient method to log and handle internal server errors.
func HandleInternalServerError(w http.ResponseWriter, r *http.Request, err error, logMsg ...string) {
	uniqueErrID := uuid.New().String()

	render.Status(r, http.StatusInternalServerError)
	render.Respond(w, r, &server.Error{
		ErrorId: uniqueErrID,
		Code: 500,
		Error: "Internal Server Error",
		ErrorType: server.InternalServerError,
		Timestamp: time.Now(),
	})
}

// HandleBadRequestError is a convenient method to log and handle bad request errors.
func HandleBadRequestError(w http.ResponseWriter, r *http.Request, err error, logMsg ...string) {
	uniqueErrID := uuid.New().String()

	render.Status(r, http.StatusBadRequest)
	render.Respond(w, r, &server.Error{
		ErrorId: uniqueErrID,
		Code: 400,
		Error: err.Error(),
		ErrorType: server.BadRequest,
		Timestamp: time.Now(),
	})
}
