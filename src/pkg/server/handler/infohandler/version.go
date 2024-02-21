package infohandler

import (
	server "linuxcode/inventory_manager/pkg/server/generated"
	"net/http"

	"github.com/go-chi/render"
)

// VersionHandler manages status related endpoints.
type VersionHandler struct {
	version     string
	buildDate   string
	description string
	commitHash  string
	commitDate  string
}

// NewVersionHandler creates a new status handler.
func NewVersionHandler(version, buildDate, description, commitHash, commitDate string) *VersionHandler {
	return &VersionHandler{
		version:     version,
		buildDate:   buildDate,
		description: description,
		commitHash:  commitHash,
		commitDate:  commitDate,
	}
}

// VersionCheck provides the linuxcode/inventory_managerlication version information.
func (h *VersionHandler) VersionCheck(w http.ResponseWriter, r *http.Request) {
	resp := server.Version{
		Version:     h.version,
		BuildDate:   h.buildDate,
		Description: h.description,
		CommitHash:  h.commitHash,
		CommitDate:  h.commitDate,
	}

	render.Respond(w, r, &resp)
}
