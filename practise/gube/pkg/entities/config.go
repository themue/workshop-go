// Go Workshop - Practise - Gube - Entities

package entities

// Config provides configuration for the runnables.
// Environment may contain multiple of them.
type Config struct {
	ID        string            `json:"id"`
	KeyValues map[string]string `json:"key_values"`
}
