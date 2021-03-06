// Go Workshop - Practise - Gube - Runner

package runner

import (
	"fmt"
	"sync"

	"github.com/themue/workshop-go/practise/gube/pkg/runnable"
)

// Notifier defines a callback interface to allow the
// Runner a callback communication with the Manager.
type Notifier interface {
	// NotifyRunnerError tells the method implementor
	// that a Runner has ended with the given error.
	NotifyRunnerError(id string, err error)

	// NotifyRunnerPanic tells the method implementor
	// that a Runner has panicked with the given error.
	NotifyRunnerPanic(id string, err error)
}

// Runner is a little helper used by the Manager. It allows
// to spawn and kill a Runnable and cares for troubles during
// runtime.
type Runner struct {
	mu          sync.Mutex
	runnable    runnable.Runnable
	environment *runnable.Environment
	notifier    Notifier
	running     bool
	err         error
}

// New creates a new Runner instance.
func New(
	run runnable.Runnable,
	env *runnable.Environment,
	notifier Notifier,
) *Runner {
	return &Runner{
		runnable:    run,
		environment: env,
		notifier:    notifier,
	}
}

// Spawn starts the Runnable.
func (r *Runner) Spawn() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.running {
		// Runner already runs.
		return nil
	}
	// Start the Runnable as goroutine.
	go r.run()
	return nil
}

// Kill tells the Runnable to stop working.
func (r *Runner) Kill() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !r.running {
		if r.err != nil {
			return fmt.Errorf("Runner does not run ID %q: %v", r.runnable.ID(), r.err)
		}
		return fmt.Errorf("Runner does not run ID %q", r.runnable.ID())
	}
	r.err = r.runnable.Stop()
	return r.err
}

// Err returns the current error status of the Runner.
func (r *Runner) Err() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.err
}

// run will be started as goroutine and makes the Runnable
// work as well as it cares of the status and listeners when
// it ends.
func (r *Runner) run() {
	var err error
	defer func() {
		// Care for status and possible listeners.
		r.mu.Lock()
		defer r.mu.Unlock()
		for _, storage := range r.environment.Storages {
			storage.RemoveListener(r.runnable.ID())
		}
		r.running = false
		r.err = err
	}()
	defer func() {
		// Care for panics.
		if rec := recover(); rec != nil {
			err = fmt.Errorf("Runner %q had panic: %v", r.runnable.ID(), rec)
			go r.notifier.NotifyRunnerPanic(r.runnable.ID(), err)
		}
	}()
	// Now let the runnable run.
	if err = r.runnable.Run(r.environment); err != nil {
		go r.notifier.NotifyRunnerError(r.runnable.ID(), err)
	}
}
