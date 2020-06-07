// Go Workshop - Practise - Gube - Registry

package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/themue/workshop-go/practise/gube/pkg/actor"
	"github.com/themue/workshop-go/practise/gube/pkg/entities"
)

// entitySets is used as a container for backup and restore in JSON.
type entitySets struct {
	Configs      map[string]entities.Config      `json:"configs"`
	Environments map[string]entities.Environment `json:"environments"`
	Storages     map[string]entities.Storage     `json:"storages"`
}

// InMemoryRegistry implements the Registry interface using a
// simple number of in-memory maps per entity. It provides concurrent
// usage.
type InMemoryRegistry struct {
	data entitySets
	act  *actor.Actor
}

// NewInMemoryRegistry creates an instance of the in-memory
// registry. The concrete type is returned to allow white-box
// testing or use other exported methods not defined in the
// Registry interface.
func NewInMemoryRegistry(ctx context.Context) *InMemoryRegistry {
	return &InMemoryRegistry{
		data: entitySets{
			Configs:      make(map[string]entities.Config),
			Environments: make(map[string]entities.Environment),
			Storages:     make(map[string]entities.Storage),
		},
		act: actor.New(actor.Context(ctx)),
	}
}

// StoreConfig implements the Registry interface.
func (r *InMemoryRegistry) StoreConfig(in entities.Config) (entities.Config, error) {
	var out entities.Config

	err := r.act.DoSync(func() {
		tmp := r.data.Configs[in.ID]
		tmp.DeepCopyInto(&out)
		r.data.Configs[in.ID] = in
	}, time.Second)

	return out, err
}

// RetrieveConfig implements the Registry interface.
func (r *InMemoryRegistry) RetrieveConfig(id string) (entities.Config, error) {
	var out entities.Config

	if err := r.act.DoSync(func() {
		tmp := r.data.Configs[id]
		tmp.DeepCopyInto(&out)
	}, time.Second); err != nil {
		return out, err
	}

	if out.ID != id {
		return out, fmt.Errorf("Config with ID %q not found", id)
	}

	return out, nil
}

// FindConfigs implements the Registry interface.
func (r *InMemoryRegistry) FindConfigs(matches func(c entities.Config) bool) ([]entities.Config, error) {
	var outs []entities.Config
	var ferr error

	if err := r.act.DoSync(func() {
		var out entities.Config
		for _, tmp := range r.data.Configs {
			tmp.DeepCopyInto(&out)
			if matches(out) {
				outs = append(outs, out)
			}
		}
	}, 2*time.Second); err != nil {
		return outs, err
	}

	return outs, ferr
}

// StoreEnvironment implements the Registry interface.
func (r *InMemoryRegistry) StoreEnvironment(in entities.Environment) (entities.Environment, error) {
	var out entities.Environment

	err := r.act.DoSync(func() {
		tmp := r.data.Environments[in.ID]
		tmp.DeepCopyInto(&out)
		r.data.Environments[in.ID] = in
	}, time.Second)

	return out, err
}

// RetrieveEnvironment implements the Registry interface.
func (r *InMemoryRegistry) RetrieveEnvironment(id string) (entities.Environment, error) {
	var out entities.Environment

	if err := r.act.DoSync(func() {
		tmp := r.data.Environments[id]
		tmp.DeepCopyInto(&out)
	}, time.Second); err != nil {
		return out, err
	}

	if out.ID != id {
		return out, fmt.Errorf("Environment with ID %q not found", id)
	}

	return out, nil
}

// StoreStorage implements the Registry interface.
func (r *InMemoryRegistry) StoreStorage(in entities.Storage) (entities.Storage, error) {
	var out entities.Storage

	err := r.act.DoSync(func() {
		tmp := r.data.Storages[in.ID]
		tmp.DeepCopyInto(&out)
		r.data.Storages[in.ID] = in
	}, time.Second)

	return out, err
}

// RetrieveStorage implements the Registry interface.
func (r *InMemoryRegistry) RetrieveStorage(id string) (entities.Storage, error) {
	var out entities.Storage

	if err := r.act.DoSync(func() {
		tmp := r.data.Storages[id]
		tmp.DeepCopyInto(&out)
	}, time.Second); err != nil {
		return out, err
	}

	if out.ID != id {
		return out, fmt.Errorf("Storage with ID %q not found", id)
	}

	return out, nil
}

// Backup allows to write all data marshalled to JSON to a Writer.
func (r *InMemoryRegistry) Backup(out io.Writer) error {
	var werr error

	if err := r.act.DoSync(func() {
		enc := json.NewEncoder(out)
		werr = enc.Encode(r.data)
	}, 10*time.Second); err != nil {
		return err
	}

	return werr
}

// Restore allows to read all data as JSON from a Reader.
func (r *InMemoryRegistry) Restore(in io.Reader) error {
	var rerr error

	if err := r.act.DoSync(func() {
		var data entitySets
		dec := json.NewDecoder(in)
		rerr = dec.Decode(&data)
		r.data = data
	}, 10*time.Second); err != nil {
		return err
	}

	return rerr
}
