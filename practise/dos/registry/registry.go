// Go Workshop - Practise - Data, Objects, and Services

package objects

import (
	"github.com/themue/workshop-go/practise/dos/data"
)

// Registry maintains a number of entities. It does not
// handle concurreny, this will be task of the according
// service.
type Registry struct {
	storages    map[string]data.Storage
	files       map[string]data.File
	executables map[string]data.Executable
}

// New creates a registry instance.
func New() *Registry {
	return &Registry{
		storages:    make(map[string]data.Storage),
		files:       make(map[string]data.File),
		executables: make(map[string]data.Executable),
	}
}
