package infohandler

import (
	"net/http"

	"github.com/go-chi/render"
)

// Just a healthcheck that returns HTTP STATUS 200 - OK
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
	render.Respond(w, r, "healthy")
}
