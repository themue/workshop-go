// Go Workshop - Practise - Data, Objects, and Services

package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/themue/workshop-go/practise/dos/data"
	"github.com/themue/workshop-go/practise/dos/registry"
)

const (
	defaultTimeout = 5 * time.Second
)

// Manager is a service able to handle environments and executables.
type Manager struct {
	ctx      context.Context
	registry registry.Registry
	runner   Runner
	actions  chan func()
	err      error
}

// New creates a new Manager service running in the background.
func New(ctx context.Context, registry registry.Registry, runner Runner) *Manager {
	m := &Manager{
		ctx:      ctx,
		registry: registry,
		runner:   runner,
		actions:  make(chan func()),
	}

	go m.backend()

	return m
}

// Apply sets a number of data objects in the registry.
func (m *Manager) Apply(ins ...interface{}) ([]interface{}, error) {
	var outs []interface{}
	var applyErr error
	err := m.do(func() {
		for i, in := range ins {
			switch d := in.(type) {
			case data.File:
				if out, ok := m.registry.SetFile(d); ok {
					outs = append(outs, out)
				}
			case data.Executable:
				if out, ok := m.registry.SetExecutable(d); ok {
					outs = append(outs, out)
				}
			case data.Storage:
				if out, ok := m.registry.SetStorage(d); ok {
					outs = append(outs, out)
				}
			default:
				applyErr = fmt.Errorf("applied %d has invalid type %T", i, d)
				return
			}
		}
	}, 2*defaultTimeout)
	if err != nil {
		return nil, err
	}
	return outs, applyErr
}

// Run starts the executable with the given ID.
func (m *Manager) Run(id string) (RunningID, error) {
	var runningID RunningID
	var runErr error
	err := m.do(func() {
		// Retrieve executable.
		exec, ok := m.registry.Executable(id)
		if !ok {
			runErr = fmt.Errorf("executable %q not found", id)
			return
		}
		// Retrieve file and file storage.
		file, ok := m.registry.File(exec.FileID)
		if !ok {
			runErr = fmt.Errorf("file %q not found", exec.FileID)
			return
		}
		fileStorage, ok := m.registry.Storage(file.StorageID)
		// Retrieve storages of the executable.
		storages := []data.Storage{fileStorage}
		for _, sid := range exec.StorageIDs {
			storage, ok := m.registry.Storage(sid)
			if !ok {
				runErr = fmt.Errorf("storage %q not found", sid)
				return
			}
			storages = append(storages, storage)
		}
		// All fine, now let it run.
		runningID, runErr = m.runner.Run(file, exec, storages...)
	}, defaultTimeout)
	if err != nil {
		return "", err
	}
	return runningID, runErr
}

// do performs an action with a timeout.
func (m *Manager) do(action func(), timeout time.Duration) error {
	select {
	case m.actions <- action:
		// All fine.
		return nil
	case <-time.After(timeout):
		// Waited for too long, possibly backend broken.
		return fmt.Errorf("timeout while waiting to perform action")
	}
}

// backend runs the Manager actions one by one.
func (m *Manager) backend() {
	defer func() {
		// Take care if an action panics.
		if r := recover(); r != nil {
			m.err = fmt.Errorf("panic while performing action: %v", r)
		}
	}()
	for {
		select {
		case <-m.ctx.Done():
			if m.err == nil {
				m.err = m.ctx.Err()
			}
			return
		case action := <-m.actions:
			action()
		}
	}
}
