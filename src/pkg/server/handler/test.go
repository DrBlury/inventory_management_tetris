package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// Just a Testhandler returns HTTP STATUS 200 - OK with a JSON response
func TestHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
	render.Respond(w, r, "Hello!")
}
