// Go Workshop - Practise - Gube - Entities

package entities

// Config provides configuration for the runnables.
// Environment may contain multiple of them.
type Config struct {
	ID        string            `json:"id"`
	KeyValues map[string]string `json:"key_values"`
}

// DeepCopyInto copies all fields into the given instance.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.KeyValues != nil {
		out.KeyValues = make(map[string]string)
		for k, v := range in.KeyValues {
			out.KeyValues[k] = v
		}
	}
}
