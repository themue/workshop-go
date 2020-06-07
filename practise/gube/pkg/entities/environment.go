// Go Workshop - Practise - Gube - Entities

package entities

// Environment contains the environment a runnable
// is started with.
type Environment struct {
	ID         string   `json:"id"`
	ConfigIDs  []string `json:"config_ids"`
	StorageIDs []string `json:"storage_ids"`
}

// DeepCopyInto copies all fields into the given instance.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	if in.ConfigIDs != nil {
		out.ConfigIDs = make([]string, len(in.ConfigIDs))
		copy(out.ConfigIDs, in.ConfigIDs)
	}
	if in.StorageIDs != nil {
		out.StorageIDs = make([]string, len(in.StorageIDs))
		copy(out.StorageIDs, in.StorageIDs)
	}
}
