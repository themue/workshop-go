// Go Workshop - Practise - Gube - Entities

package entities

// StorageType descibes different types of storage.
type StorageType string

// Definition of different storrage types.
const (
	InternalStorageType StorageType = "internal"
	RedisStorageType                = "redis"
)

// Storage defines a Key/Value storage instance.
type Storage struct {
	ID         string            `json:"id"`
	Type       StorageType       `json:"type"`
	TypeConfig map[string]string `json:"type_config"`
}

// DeepCopyInto copies all fields into the given instance.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.TypeConfig != nil {
		out.TypeConfig = make(map[string]string)
		for k, v := range in.TypeConfig {
			out.TypeConfig[k] = v
		}
	}
}
