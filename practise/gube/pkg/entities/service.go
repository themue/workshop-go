// Go Workshop - Practise - Gube - Entities

package entities

// Service defines the association between a deployed
// Runnable and its Environment.
type Service struct {
	ID            string `json:"id"`
	EnvironmentID string `json:"environment_id"`
	Restart       bool   `json:"restart"`
}

// DeepCopyInto copies all fields into the given instance.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}
