// Go Workshop - Practise - Gube - Manager

package manager

import (
	"fmt"
	"net/http"

	"github.com/themue/workshop-go/practise/gube/internal/handler"
)

// api cares for the RESTful API.
type api struct {
	manager *Manager
	server  http.Server
}

// newAPI starts a server on the given port with the
// Gube handlers.
func newAPI(m *Manager, port int) *api {
	mux := http.NewServeMux()

	mux.Handle("/configs", handler.NewConfigsHandler(m.registry))

	a := &api{
		manager: m,
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}

	go func() {
		err := a.server.ListenAndServe()
		a.manager.logf("API server terminated with error: %v", err)
	}()

	return a
}
