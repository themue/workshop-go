// Go Workshop - Practise - Data, Objects, and Services

package manager

import (
	"github.com/themue/workshop-go/practise/dos/data"
)

// RunningID is returned by Runner.Run to identify a spawned
// executable. So it later can be managed.
type RunningID string

// Runner defines the interface for types able to run
// an executable. Implementations could be e.g. for Docker.
type Runner interface {
	// Run starts the given file as defined executable with
	// storages. It returnes an ID for later management.
	Run(file data.File, exec data.Executable, storages ...data.Storage) (RunningID, error)

	// Kill terminates the executable with the given ID.
	Kill(id RunningID) error
}

// NewDefaultRunner returns a default runner implementation.
func NewDefaultRunner() Runner {
	return nil
}
