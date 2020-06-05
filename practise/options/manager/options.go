// Go Workshop - Practise - Options

package manager

import (
	"context"

	"github.com/themue/workshop-go/practise/dos/manager"
	"github.com/themue/workshop-go/practise/dos/registry"
)

// Option is a function able to modify a Manager.
type Option func(m *Manager)

// Context sets the Manager context. Default is the
// background context.
func Context(ctx context.Context) Option {
	return func(m *Manager) {
		if ctx != nil {
			m.ctx = ctx
		} else {
			m.ctx = context.Background()
		}
	}
}

// Registry sets the Manager registry. Default is the
// simple registry.
func Registry(registry registry.Registry) Option {
	return func(m *Manager) {
		if registry != nil {
			m.registry = registry
		} else {
			m.registry = registry.NewSimpleRegistry()
		}
	}
}

// Runner sets the Manager runner. Default is the
// default runner.
func Runner(runner manager.Runner) Option {
	return func(m *Manager) {
		if runner != nil {
			m.runner = runner
		} else {
			m.runner = manager.NewDefaultRunner()
		}
	}
}
