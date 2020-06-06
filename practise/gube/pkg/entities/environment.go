// Go Workshop - Practise - Gube - Entities

package entities

// Environment contains the environment a runnable
// is started with.
type Environment struct {
	ID         string   `json:"id"`
	ConfigIDs  []string `json:"config_ids"`
	StorageIDs []string `json:"storage_ids"`
}
