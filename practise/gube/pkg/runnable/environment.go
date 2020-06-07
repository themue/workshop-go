// Go Workshop - Practise - Gube - Runnable

package runnable

import "github.com/themue/workshop-go/practise/gube/pkg/storage"

// Environment contains the environment passed to a
// Runnable for work.
type Environment struct {
	Configurations map[string]Config
	Storages       map[string]storage.Storage
}
