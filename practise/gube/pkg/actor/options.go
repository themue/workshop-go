// Go Workshop - Practise - Gube - Actor

package actor

import (
	"context"
)

// Option is a function able to modify an Acor.
type Option func(act *Actor)

// Context sets the Actor context. Default is the
// background context.
func Context(ctx context.Context) Option {
	return func(act *Actor) {
		if ctx != nil {
			act.ctx = ctx
		}
	}
}

// Buffersize sets the buffer size of the Actor channel
// for asynchronous actions. Minimum is 16.
func Buffersize(size int) Option {
	return func(act *Actor) {
		if size > 16 {
			act.asyncs = make(chan func(), size)
		}
	}
}
