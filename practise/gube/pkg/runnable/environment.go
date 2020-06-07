// Go Workshop - Practise - Gube - Runnable

package runnable

import (
	"github.com/themue/workshop-go/practise/gube/pkg/storage"
)

// Logger allows a runnable to use the central logger.
type Logger interface {
	Logf(format string, v ...interface{})
}

// Environment contains the environment passed to a
// Runnable for work.
type Environment struct {
	Logger         Logger
	Configurations map[string]Config
	Storages       map[string]storage.Storage
}
