// Go Workshop - Practise - Gube - Manager

// Package manager provides the main type Manager for
// interaction with Gube.
package manager

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/themue/workshop-go/practise/gube/internal/runner"
	"github.com/themue/workshop-go/practise/gube/pkg/actor"
	"github.com/themue/workshop-go/practise/gube/pkg/entities"
	"github.com/themue/workshop-go/practise/gube/pkg/registry"
	"github.com/themue/workshop-go/practise/gube/pkg/runnable"
	"github.com/themue/workshop-go/practise/gube/pkg/storage"
)

// Manager is the one central manager for Gube. Here
// the different entities and runnables are deployed
// and services started or stopped.
type Manager struct {
	mu        sync.RWMutex
	ctx       context.Context
	registry  registry.Registry
	storages  map[string]storage.Storage
	runnables map[string]runnable.Runnable
	runners   map[string]*runner.Runner
	act       *actor.Actor
	logger    *log.Logger
	callback  *callback
	err       error
}

// New creates a configured Gube Manager instance.
func New(options ...Option) *Manager {
	m := &Manager{
		ctx:       context.Background(),
		storages:  make(map[string]storage.Storage),
		runnables: make(map[string]runnable.Runnable),
		runners:   make(map[string]*runner.Runner),
		logger:    log.New(os.Stdout, "gube", log.Ldate|log.Ltime|log.Lshortfile),
	}
	m.callback := &callback{
		manager: m,
	}

	for _, option := range options {
		option(m)
	}

	// Post-precessing of configuration.
	m.act = actor.New((actor.Context(m.ctx)))
	if m.registry == nil {
		m.registry = registry.NewInMemoryRegistry(m.ctx)
	}

	return m
}

// Deploy adds a number of entities plus runnables
// to the Manager.
func (m *Manager) Deploy(ents ...interface{}) {
	if err := m.act.DoAsync(func() {
		for _, raw := range ents {
			var err error
			switch entity := raw.(type) {
			case entities.Config:
				_, err = m.registry.StoreConfig(entity)
			case entities.Environment:
				_, err = m.registry.StoreEnvironment(entity)
			case entities.Service:
				_, err = m.registry.StoreService(entity)
			case entities.Storage:
				_, err = m.registry.StoreStorage(entity)
				if err != nil {
					switch entity.Type {
					case entities.InternalStorageType:
						m.storages[entity.ID] = storage.NewInMemoryStorage(m.ctx)
					default:
						err = fmt.Errorf("storage type %q not yet implemented", entity.Type)
					}
				}
			case runnable.Runnable:
				m.runnables[entity.ID()] = entity
			default:
				err = fmt.Errorf("invalid entity type %T", raw)
			}
			// Check potential error.
			if err != nil {
				m.Logf("error during deployment: %v", err)
			}
		}
	}, 5*time.Second); err != nil {
		m.err = err
		m.Logf("deployment of entities failed")
	}
}

// Spawn starts the Service with the given ID.
func (m *Manager) Spawn(id string) error {
	// Check already spawned (and potentially killed) Runners first.
	run, ok := m.runners[id]
	if !ok {
		// Prepare a new Runner.
		var err error
		run, err = m.prepareRunner(id)
		if err != nil {
			return err
		}
	}
	// Now spawn the Runner.
	return run.Spawn()
}

// Err returns the current error status.
func (m *Manager) Err() error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.err
}

// prepareRunner creates a Runner.
func (m *Manager) prepareRunner(id string) (*runner.Runner, error) {
	svc, err := m.registry.RetrieveService(id)
	if err != nil {
		return nil, err
	}
	runable := m.runnables[svc.ID]
	if runable == nil {
		return nil, fmt.Errorf("runnable %q not found", svc.ID)
	}
	env, err := m.prepareEnvironment(svc.EnvironmentID)
	if err != nil {
		return nil, err
	}
	run := runner.New(runable, env, m.callback)
	return run, nil
}

// prepareEnvironment creates a runnable Environment.
func (m *Manager) prepareEnvironment(id string) (*runnable.Environment, error) {
	env := &runnable.Environment{
		Logger:         m.callback,
		Configurations: make(map[string]runnable.Config),
		Storages:       make(map[string]storage.Storage),
	}
	eenv, err := m.registry.RetrieveEnvironment(id)
	if err != nil {
		return nil, err
	}
	// Copy configuration values.
	for _, id := range eenv.ConfigIDs {
		econf, err := m.registry.RetrieveConfig(id)
		if err != nil {
			return nil, err
		}
		rconf := runnable.Config{
			Values: make(map[string]string),
		}
		for k, v := range econf.KeyValues {
			rconf.Values[k] = v
		}
		env.Configurations[id] = rconf
	}
	// Copy storages.
	for _, id := range eenv.StorageIDs {
		store := m.storages[id]
		if store == nil {
			return nil, fmt.Errorf("storage %q of environment not found", id)
		}
		env.Storages[id] = store
	}
	return env, nil
}
