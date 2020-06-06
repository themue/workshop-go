// Go Workshop - Language - White Testing

package calculator

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatalf("New returned nil")
	}
	if c.registers == nil {
		t.Fatalf("registers in Calculator is nil")
	}
	if len(c.registers) != 0 {
		t.Fatalf("registers in Calculator are not empty")
	}
}

func TestAdd(t *testing.T) {
	c := New()

	c.Add("a", 1.0, 2.0, 3.0)
	c.Add("b", 5.0, -5.0)
	c.Add("c")

	if len(c.registers) != 2 {
		t.Errorf("wrong number of registers: %d", len(c.registers))
	}
	if c.registers["a"] != 6.0 {
		t.Errorf("illegal value of register 'a': %f", c.registers["a"])
	}
	if c.registers["b"] != 0.0 {
		t.Errorf("illegal value of register 'b': %f", c.registers["b"])
	}
	if c.registers["c"] != 0.0 {
		t.Errorf("illegal value of register 'c': %f", c.registers["c"])
	}
}
