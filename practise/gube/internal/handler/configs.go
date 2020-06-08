// Go Workshop - Practise - Gube - Handler

package handler

import (
	"net/http"

	"github.com/themue/workshop-go/practise/gube/internal/httpx"
	"github.com/themue/workshop-go/practise/gube/pkg/entities"
	"github.com/themue/workshop-go/practise/gube/pkg/registry"
)

// ConfigsHandler cares for RESTful requests for Configs.
type ConfigsHandler struct {
	registry registry.Registry
}

// NewConfigHandler ...
func NewConfigHandler(registry registry.Registry) *ConfigsHandler {
	return &ConfigsHandler{
		registry: registry,
	}
}

// ServeHTTPGet reads Configs out of registry.
func (ch *ConfigsHandler) ServeHTTPGet(w http.ResponseWriter, r *http.Request) {
	parts := httpx.PathParts(r)
	switch len(parts) {
	case 1:
		// Request GET /configs wants to read all Configs.
		http.Error(w, "getting all Configs not yet implemented", http.StatusNotImplemented)
	case 2:
		// Request GET /configs/<id> wants to read one Config.
		cfg, err := ch.registry.RetrieveConfig(parts[1])
		if err != nil {
			http.Error(w, "Config not found", http.StatusNotFound)
			return
		}
		w.Header().Add(httpx.HeaderContentType, httpx.ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		httpx.MarshalBody(w, w.Header(), cfg)
	default:
		// Request too long.
		http.Error(w, "request too long", http.StatusBadRequest)
	}
}

// ServeHTTPPost writes Configs into the registry.
func (ch *ConfigsHandler) ServeHTTPPost(w http.ResponseWriter, r *http.Request) {
	parts := httpx.PathParts(r)
	switch len(parts) {
	case 1:
		// Request POST /configs wants to write many Configs.
		http.Error(w, "writing many Configs not yet implemented", http.StatusNotImplemented)
	case 2:
		// Request POST /configs/<id> wants to write one Config.
		var cfg entities.Config
		err := httpx.UnmarshalBody(r.Body, r.Header, &cfg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if parts[1] != cfg.ID {
			http.Error(w, "IDs doesn't match", http.StatusExpectationFailed)
			return
		}
		out, err := ch.registry.StoreConfig(cfg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add(httpx.HeaderContentType, httpx.ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		httpx.MarshalBody(w, w.Header(), out)
	default:
		// Request too long.
		http.Error(w, "request too long", http.StatusBadRequest)
	}
}

// ServeHTTP implements http.Handler.
func (ch *ConfigsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method for Config handler not allowed", http.StatusMethodNotAllowed)
}
