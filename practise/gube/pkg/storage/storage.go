// Go Workshop - Practise - Gube - Storage

package storage

// Storage defines the simple functionality of key/value
// databases in memory or externally.
type Storage interface {
	// Put writes a key/value combination. An existing value
	// will be returned.
	Put(key, value string) (string, error)

	// Get reads a value out of the storage.
	Get(key string) (string, error)

	// Find returns all key/values where matches returns true.
	Find(matches func(key, value string) bool) (map[string]string, error)
}
