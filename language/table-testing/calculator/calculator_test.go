// Go Workshop - Language - Table-Driven Testing

package calculator_test

import (
	"testing"

	"tideland.dev/go/audit/asserts"

	"github.com/themue/workshop-go/language/table-testing/calculator"
)

func TestAdd(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)
	tests := []struct {
		title       string
		adds        map[string][]float64
		values      map[string]float64
		registerLen int
	}{
		{
			title:       "simple initialization",
			registerLen: 0,
		},
		{
			title: "one register, zero values",
			adds: map[string][]float64{
				"a": {},
			},
			values: map[string]float64{
				"a": 0.0,
			},
			registerLen: 0,
		},
		{
			title: "one register, multiple values",
			adds: map[string][]float64{
				"a": {1.0, 2.0, 3.0},
			},
			values: map[string]float64{
				"a": 6.0,
			},
			registerLen: 1,
		},
		{
			title: "two register, multiple values",
			adds: map[string][]float64{
				"a": {1.0, 2.0, 3.0},
				"b": {-5.0, -5.0},
			},
			values: map[string]float64{
				"a": 6.0,
				"b": -10.0,
			},
			registerLen: 2,
		},
	}

	for i, test := range tests {
		assert.Logf("test #%d: %s", i, test.title)
		c := calculator.New()
		assert.NotNil(c, "New returned nil")
		assert.Empty(c.RegisterKeys(), "initial calculator registers aren't empty")

		for register, addValues := range test.adds {
			c.Add(register, addValues...)
		}
		for register, value := range test.values {
			assert.Equal(c.Register(register), value)
		}
		assert.Length(c.RegisterKeys(), test.registerLen)
	}
}
