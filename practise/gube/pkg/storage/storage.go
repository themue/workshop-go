// Go Workshop - Practise - Gube - Storage

package storage

// Operations for notification of a listener.
const (
	PutOp    = "put"
	DeleteOp = "delete"
)

// Listener allows Runnables to listen to changes to
// a storage. The Listener will be called
type Listener func(op, key, value string)

// Storage defines the simple functionality of key/value
// databases in memory or externally.
type Storage interface {
	// Put writes a key/value combination. An existing value
	// will be returned.
	Put(key, value string) (string, error)

	// Get reads a value out of the storage.
	Get(key string) (string, error)

	// Delete removes a value from the storage.
	Delete(key string) (string, error)

	// Find returns all key/values where matches returns true.
	Find(matches func(key, value string) bool) (map[string]string, error)

	// AddListener adds a listener to the storage.
	AddListener(id string, l Listener)

	// RemoveListener removes a listener from the storage.
	RemoveListener(id string)
}
