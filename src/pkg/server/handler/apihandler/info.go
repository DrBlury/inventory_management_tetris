package apihandler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Get metrics
// (GET /info/metrics)
func (a APIHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

// Get status
// (GET /info/status)
func (a APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, &struct{ Status string }{"HEALTHY"})
}

// Get version
// (GET /info/version)
func (a APIHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, a.Info)
}
