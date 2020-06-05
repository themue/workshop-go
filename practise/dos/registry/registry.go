// Go Workshop - Practise - Data, Objects, and Services

package registry

import (
	"github.com/themue/workshop-go/practise/dos/data"
)

// Registry defines the functionality any implementation
// shall provide.
type Registry interface {
	// SetStorage puts a storage into the registry and returns
	// a possible existing one with the same ID.
	SetStorage(in data.Storage) (data.Storage, bool)

	// Storage returns the storage with the given ID.
	Storage(id string) (data.Storage, bool)

	// FindStorage returns all storages where the function matches
	// returns true.
	FindStorage(matches func(s data.Storage) bool) []data.Storage

	// SetFile puts a file into the registry and returns
	// a possible existing one with the same ID.
	SetFile(in data.File) (data.File, bool)

	// File returns the file with the given ID.
	File(id string) (data.File, bool)

	// SetExecutable puts an executable into the registry and returns
	// a possible existing one with the same ID.
	SetExecutable(in data.Executable) (data.Executable, bool)

	// Executable returns the executable with the given ID.
	Executable(id string) (data.Executable, bool)
}

// SimpleRegistry maintains a number of entities. It does not
// handle concurreny, this will be task of the according
// service.
type SimpleRegistry struct {
	storages    map[string]data.Storage
	files       map[string]data.File
	executables map[string]data.Executable
}

// NewSimpleRegistry creates a SimpleRegistry instance.
func NewSimpleRegistry() Registry {
	return &SimpleRegistry{
		storages:    make(map[string]data.Storage),
		files:       make(map[string]data.File),
		executables: make(map[string]data.Executable),
	}
}

// SetStorage implements the Registry interface.
func (r *SimpleRegistry) SetStorage(in data.Storage) (data.Storage, bool) {
	out, ok := r.storages[in.ID]
	r.storages[in.ID] = in
	return out, ok
}

// Storage implements the Registry interface.
func (r *SimpleRegistry) Storage(id string) (data.Storage, bool) {
	out, ok := r.storages[id]
	return out, ok
}

// FindStorage implements the Registry interface.
func (r *SimpleRegistry) FindStorage(matches func(s data.Storage) bool) []data.Storage {
	var found []data.Storage
	for _, s := range r.storages {
		if matches(s) {
			found = append(found, s)
		}
	}
	return found
}

// SetFile implements the Registry interface.
func (r *SimpleRegistry) SetFile(in data.File) (data.File, bool) {
	out, ok := r.files[in.ID]
	r.files[in.ID] = in
	return out, ok
}

// File implements the Registry interface.
func (r *SimpleRegistry) File(id string) (data.File, bool) {
	out, ok := r.files[id]
	return out, ok
}

// SetExecutable implements the Registry interface.
func (r *SimpleRegistry) SetExecutable(in data.Executable) (data.Executable, bool) {
	out, ok := r.executables[in.ID]
	r.executables[in.ID] = in
	return out, ok
}

// Executable implements the Registry interface.
func (r *SimpleRegistry) Executable(id string) (data.Executable, bool) {
	out, ok := r.executables[id]
	return out, ok
}
