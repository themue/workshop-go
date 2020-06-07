// Go Workshop - Practise - Gube - Storage

package storage

// Storage defines the simple functionality of key/value
// databases in memory or externally.
type Storage interface {
	// Put writes a key/value combination. An existing value
	// will be returned.
	Put(in string) (string, error)

	// Get reads a value out of the storage.
	Get(key string) (string, error)

	// Find returns all key/values where matches returns true.
	FindConfigs(matches func(key, value string) bool) (map[string]string, error)
}
