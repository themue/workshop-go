// Go Workshop - Practise - Gube - Registry

package registry

import (
	"github.com/themue/workshop-go/practise/gube/pkg/entities"
)

// Registry defines the functionality different entity store
// implementations (in-memory, persistent) have to provide.
type Registry interface {
	// StoreConfig writes a Config into the registry. A possible
	// existing one will be returned.
	StoreConfig(in entities.Config) (entities.Config, error)

	// RetrieveConfig reads a Config out of the registry.
	RetrieveConfig(id string) (entities.Config, error)

	// FindConfigs returns all where matches returns true.
	FindConfigs(matches func(c entities.Config) bool) ([]entities.Config, error)

	// StoreEnvironment writes an Environment into the registry. A possible
	// existing one will be returned.
	StoreEnvironment(in entities.Environment) (entities.Environment, error)

	// RetrieveEnvironment reads an Environment out of the registry.
	RetrieveEnvironment(id string) (entities.Environment, error)

	// StoreService writes a Service into the registry. A possible
	// existing one will be returned.
	StoreService(in entities.Service) (entities.Service, error)

	// RetrieveService reads a Service out of the registry.
	RetrieveService(id string) (entities.Service, error)

	// StoreStorage writes a Storage into the registry. A possible
	// existing one will be returned.
	StoreStorage(in entities.Storage) (entities.Storage, error)

	// RetrieveStorage reads a Storage out of the registry.
	RetrieveStorage(id string) (entities.Storage, error)
}
