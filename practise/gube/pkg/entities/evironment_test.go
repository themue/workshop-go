// Go Workshop - Practise - Gube - Entities

package entities_test

import (
	"testing"

	"github.com/themue/workshop-go/practise/gube/pkg/entities"
)

func TestEnvironment(t *testing.T) {
	in := entities.Environment{
		ID: "environment",
		ConfigIDs: []string{
			"a", "b", "c",
		},
		StorageIDs: []string{
			"d", "e", "f",
		},
	}

	var out entities.Environment
	in.DeepCopyInto(&out)

	if out.ID != in.ID {
		t.Errorf("Environment: ID not copied")
	}
	if len(out.ConfigIDs) != len(in.ConfigIDs) {
		t.Errorf("Environment: ConfigIDs has wrong length")
	}
	for i, id := range out.ConfigIDs {
		if id != in.ConfigIDs[i] {
			t.Errorf("Environment: ConfigIDs not copied correctly")
		}
	}
}
