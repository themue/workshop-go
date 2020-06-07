// Go Workshop - Practise - Gube - Handler

package handler

import (
	"net/http"

	"github.com/themue/workshop-go/practise/gube/pkg/registry"
)

// ConfigHandler cares for RESTful requests for Config
type ConfigHandler struct {
	registry registry.Registry
}

// NewConfigHandler ...
func NewConfigHandler(registry registry.Registry) *ConfigHandler {
	return &ConfigHandler{
		registry: registry,
	}
}

// ServeHTTPGet reads a Config out of registry.
func (ch *ConfigHandler) ServeHTTPGet(w http.ResponseWriter, r *http.Request) {

}

// ServeHTTP implements http.Handler.
func (ch *ConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
