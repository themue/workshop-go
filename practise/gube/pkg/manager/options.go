// Go Workshop - Practise - Gube - Manager

package manager

import (
	"context"

	"github.com/themue/workshop-go/practise/gube/pkg/registry"
)

// Option is a function able to modify a Manager.
type Option func(m *Manager)

// Context sets the Manager context. Default is the
// background context.
func Context(ctx context.Context) Option {
	return func(m *Manager) {
		if ctx != nil {
			m.ctx = ctx
		}
	}
}

// Registry sets the Manager registry. Default is the
// in-memory registry.
func Registry(r registry.Registry) Option {
	return func(m *Manager) {
		if r != nil {
			m.registry = r
		}
	}
}
