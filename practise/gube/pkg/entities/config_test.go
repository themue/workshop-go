// Go Workshop - Practise - Gube - Entities

package entities_test

import (
	"testing"

	"github.com/themue/workshop-go/practise/gube/pkg/entities"
)

func TestConfig(t *testing.T) {
	in := entities.Config{
		ID: "config",
		KeyValues: map[string]string{
			"a": "123",
			"b": "456",
		},
	}

	var out entities.Config
	in.DeepCopyInto(&out)

	if out.ID != in.ID {
		t.Errorf("Confg: ID not copied")
	}
	if len(out.KeyValues) != len(in.KeyValues) {
		t.Errorf("Config: KeyValues has wrong length")
	}
	for k, v := range out.KeyValues {
		if v != in.KeyValues[k] {
			t.Errorf("Config: KeyValues not copied correctly")
		}
	}
}
