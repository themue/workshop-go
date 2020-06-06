// Go Workshop - Language - Table-Driven Testing

package calculator

// Calculator is a dummy calculator for testing.
type Calculator struct {
	registers map[string]float64
}

// New creates a new calculator.
func New() *Calculator {
	return &Calculator{
		registers: make(map[string]float64),
	}
}

// Add adds the given values to the given register.
func (c *Calculator) Add(register string, values ...float64) {
	for _, value := range values {
		c.registers[register] += value
	}
}

// Sub subtracts the given values from the given register.
func (c *Calculator) Sub(register string, values ...float64) {
	for _, value := range values {
		c.registers[register] -= value
	}
}

// RegisterKeys returns the current register keys.
func (c *Calculator) RegisterKeys() []string {
	var keys []string
	for key := range c.registers {
		keys = append(keys, key)
	}
	return keys
}

// Register returns the value of the given register.
func (c *Calculator) Register(register string) float64 {
	return c.registers[register]
}

// Clear resets the given register.
func (c *Calculator) Clear(register string) {
	delete(c.registers, register)
}

// ClearAll resets all registers.
func (c *Calculator) ClearAll() {
	c.registers = make(map[string]float64)
}
