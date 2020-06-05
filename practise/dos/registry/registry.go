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

// SetStorage puts a storage into the registry and returns
// a possible existing one with the same ID.
func (r *Registry) SetStorage(in data.Storage) (data.Storage, bool) {
	out, ok := r.storages[in.ID]
	r.storages[in.ID] = in
	return out, ok
}

// Storage returns the storage with the given ID.
func (r *Registry) Storage(id string) (data.Storage, bool) {
	return r.storages[id]
}

// FindStorage returns all storages where the function matches
// returns true.
func (r *Registry) FindStorage(matches func(s data.Storage) bool) []data.Storage {
	found := []data.Storage{}
	for _, s := range r.storages {
		if matches(s) {
			found = append(found, s)
		}
	}
	return found
}
