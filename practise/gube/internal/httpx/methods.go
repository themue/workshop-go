// Go Workshop - Practise - Gube - HTTPX

package httpx

import (
	"net/http"
)

// GetHandler has to be implemented by a handler for HEAD requests
// dispatched through the MethodHandler.
type GetHandler interface {
	ServeHTTPGet(w http.ResponseWriter, r *http.Request)
}

// HeadHandler has to be implemented by a handler for HEAD requests
// dispatched through the MethodHandler.
type HeadHandler interface {
	ServeHTTPHead(w http.ResponseWriter, r *http.Request)
}

// PostHandler has to be implemented by a handler for POST requests
// dispatched through the MethodHandler.
type PostHandler interface {
	ServeHTTPPost(w http.ResponseWriter, r *http.Request)
}

// PutHandler has to be implemented by a handler for PUT requests
// dispatched through the MethodHandler.
type PutHandler interface {
	ServeHTTPPut(w http.ResponseWriter, r *http.Request)
}

// PatchHandler has to be implemented by a handler for PATCH requests
// dispatched through the MethodHandler.
type PatchHandler interface {
	ServeHTTPPatch(w http.ResponseWriter, r *http.Request)
}

// DeleteHandler has to be implemented by a handler for DELETE requests
// dispatched through the MethodHandler.
type DeleteHandler interface {
	ServeHTTPDelete(w http.ResponseWriter, r *http.Request)
}

// ConnectHandler has to be implemented by a handler for CONNECT requests
// dispatched through the MethodHandler.
type ConnectHandler interface {
	ServeHTTPConnect(w http.ResponseWriter, r *http.Request)
}

// OptionsHandler has to be implemented by a handler for OPTIONS requests
// dispatched through the MethodHandler.
type OptionsHandler interface {
	ServeHTTPOptions(w http.ResponseWriter, r *http.Request)
}

// TraceHandler has to be implemented by a handler for TRACE requests
// dispatched through the MethodHandler.
type TraceHandler interface {
	ServeHTTPTrace(w http.ResponseWriter, r *http.Request)
}

// MethodHandler checks if internal handler matches the interface
// for the HTTP method.
type MethodHandler struct {
	handler http.Handler
}

// NewMethodHandler creates a meta HTTP method handler.
func NewMethodHandler(handler http.Handler) *MethodHandler {
	if handler == nil {
		panic("need handler")
	}
	return &MethodHandler{
		handler: handler,
	}
}

// ServeHTTP implements http.Handler.
func (mh *MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if h, ok := mh.handler.(GetHandler); ok {
			h.ServeHTTPGet(w, r)
			return
		}
	case http.MethodHead:
		if h, ok := mh.handler.(HeadHandler); ok {
			h.ServeHTTPHead(w, r)
			return
		}
	case http.MethodPost:
		if h, ok := mh.handler.(PostHandler); ok {
			h.ServeHTTPPost(w, r)
			return
		}
	case http.MethodPut:
		if h, ok := mh.handler.(PutHandler); ok {
			h.ServeHTTPPut(w, r)
			return
		}
	case http.MethodPatch:
		if h, ok := mh.handler.(PatchHandler); ok {
			h.ServeHTTPPatch(w, r)
			return
		}
	case http.MethodDelete:
		if h, ok := mh.handler.(DeleteHandler); ok {
			h.ServeHTTPDelete(w, r)
			return
		}
	case http.MethodConnect:
		if h, ok := mh.handler.(ConnectHandler); ok {
			h.ServeHTTPConnect(w, r)
			return
		}
	case http.MethodOptions:
		if h, ok := mh.handler.(OptionsHandler); ok {
			h.ServeHTTPOptions(w, r)
			return
		}
	case http.MethodTrace:
		if h, ok := mh.handler.(TraceHandler); ok {
			h.ServeHTTPTrace(w, r)
			return
		}
	}
	mh.handler.ServeHTTP(w, r)
}
