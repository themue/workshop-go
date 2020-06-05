// Go Workshop - Practise - Options

package manager

import (
	"context"

	"github.com/themue/workshop-go/practise/dos/manager"
	"github.com/themue/workshop-go/practise/dos/registry"
)

// Runner as a dummy type like in DOS example.
type Runner interface{}

// DefaultRunner simulates a Runner implementation.
type DefaultRunner struct{}

// Manager as a dummy type like in DOS example.
type Manager struct {
	ctx      context.Context
	registry registry.Registry
	runner   manager.Runner
	actions  chan func()
	err      error
}

// New shows how to initialize a Manager with options.
func New(options ...Option) *Manager {
	m := &Manager{
		actions: make(chan func()),
	}
	for _, option := range options {
		option(m)
	}

	// ...

	return m
}
