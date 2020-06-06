// Go Workshop - Language - Black Box Testing

package calculator_test

import (
	"testing"

	"github.com/themue/workshop-go/language/black-testing/calculator"
)

func TestNew(t *testing.T) {
	c := calculator.New()
	if c == nil {
		t.Fatalf("New returned nil")
	}
	keys := c.RegisterKeys()
	if len(keys) != 0 {
		t.Fatalf("registers in Calculator are not empty")
	}
}

func TestAdd(t *testing.T) {
	c := calculator.New()

	c.Add("a", 1.0, 2.0, 3.0)
	c.Add("b", 5.0, -5.0)
	c.Add("c")

	keys := c.RegisterKeys()
	if len(keys) != 2 {
		t.Errorf("wrong number of registers: %d", len(keys))
	}
	if c.Register("a") != 6.0 {
		t.Errorf("illegal value of register 'a': %f", c.Register("a"))
	}
	if c.Register("b") != 0.0 {
		t.Errorf("illegal value of register 'b': %f", c.Register("b"))
	}
	if c.Register("c") != 0.0 {
		t.Errorf("illegal value of register 'c': %f", c.Register("c"))
	}
}
