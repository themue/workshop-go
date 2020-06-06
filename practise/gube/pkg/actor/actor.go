// Go Workshop - Practise - Gube - Actor

package actor

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Actor is a helping type for goroutines. It allows synchronous
// and asynchronous actions and configuring context and buffer size.
type Actor struct {
	mu     sync.RWMutex
	ctx    context.Context
	syncs  chan func()
	asyncs chan func()
	err    error
}

// New creates a new Actor to be used by services and components.
func New(options ...Option) *Actor {
	act := &Actor{
		ctx:    context.Background(),
		syncs:  make(chan func()),
		asyncs: make(chan func(), 16),
	}

	for _, option := range options {
		option(act)
	}

	go act.backend()

	return act
}

// DoSync performs a synchronous action with a timeout. When
// timeout happens the action isn't performed because the current
// action lasts to long or blocks the backend.
func (act *Actor) DoSync(action func(), timeout time.Duration) error {
	select {
	case act.syncs <- action:
		// All fine.
		return nil
	case <-time.After(timeout):
		// Waited for too long.
		return fmt.Errorf("timeout while waiting to perform action")
	}
}

// Err returns the current error status of the actor.
func (act *Actor) Err() error {
	act.mu.RLock()
	defer act.mu.RUnlock()
	return act.err
}

// DoAsync performs an asynchronous action with a timeout. When this
// timeout happens the buffer may be full and no action performed
// anymore.
func (act *Actor) DoAsync(action func(), timeout time.Duration) error {
	select {
	case act.syncs <- action:
		// All fine.
		return nil
	case <-time.After(timeout):
		// Waited for too long.
		return fmt.Errorf("timeout while waiting to perform action")
	}
}

// setErr sets the error in a concurrent save way and when it
// isn't already set.
func (act *Actor) setErr(err error) {
	act.mu.Lock()
	defer act.mu.Unlock()
	if act.err == nil {
		act.err = err
	}
}

// backend runs the Actor actions one by one.
func (act *Actor) backend() {
	defer func() {
		// Take care if an action panics.
		if r := recover(); r != nil {
			act.setErr(fmt.Errorf("panic while performing action: %v", r))
		}
	}()
	for {
		select {
		case <-act.ctx.Done():
			// Context sets we're done.
			act.setErr(act.ctx.Err())
			return
		case action := <-act.syncs:
			// Perform action via synchronous channel.
			action()
		case action := <-act.asyncs:
			// Perform action via asynchronous channel.
			action()
		}
	}
}
