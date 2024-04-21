package apihandler

import (
	_ "embed"
	"html/template"
	server "linuxcode/inventory_manager/pkg/server/generated"
	"net/http"

	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// embed the openapi JSON and HTML file into the binary
// so we can serve them without reading from the filesystem

//go:embed embedded/openapi.json
var openapiJSON []byte

//go:embed embedded/scalar.html
var openapiHTMLScalar []byte

//go:embed embedded/spotlight.html
var openapiHTMLSpotlight []byte

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

// Get openapi JSON
// (GET /info/openapi.json)
func (a APIHandler) GetOpenAPIJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(openapiJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get openapi HTML
// (GET /info/openapi.html)
func (a APIHandler) GetOpenAPIHTML(w http.ResponseWriter, r *http.Request, queryParams server.GetOpenAPIHTMLParams) {
	w.Header().Set("Content-Type", "text/html")

	var renderer string
	if queryParams.Render != nil {
		renderer = *queryParams.Render
	}

	var templateString string
	switch renderer {
	case "scalar":
		templateString = string(openapiHTMLScalar)
	case "spotlight":
		templateString = string(openapiHTMLSpotlight)
	default:
		templateString = string(openapiHTMLSpotlight)
	}
	t, err := template.New("openapi").Parse(templateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// replace the base URL in the HTML file
	// with the actual base URL of the server
	// and render to the response writer
	err = t.Execute(w, map[string]string{
		"BaseURL": a.BaseURL,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
