// Go Workshop - Practise - Data, Objects, and Services

package data

// StorageType descibes different types of storrage.
type StorageType string

// Definition of different storrage types.
const (
	FileSystemStorageType StorageType = "FS"
	KVDBStorageType                   = "KVDB"
)

// Storage describes storage like file systems to be
// used by executables.
type Storage struct {
	ID         string            `json:"id"`
	Location   string            `json:"location"`
	Type       StorageType       `json:"storage_type"`
	Parameters map[string]string `json:"parameters"`
}

// File describes a file in a directory. All upper-case
// fields are exported. The JSON tags allow to rename or
// surpress fields when marshalling to JSON. Same for XML
// and other encodings.
type File struct {
	ID         string `json:"id"`
	StorrageID string `json:"storrage_id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
}

// Executable references a file to be executed, together
// with environment variables.
type Executable struct {
	ID          string            `json:"id"`
	FileID      string            `json:"file_id"`
	Parameters  map[string]string `json:"parameters"`
	Environment map[string]string `json:"environment"`
	StorageIDs  []string          `json:"storage_ids"`
}
