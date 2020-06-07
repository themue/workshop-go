// Go Workshop - Practise - Gube - Runnable

package runnable

// Runnable defines what runnable Gube components
// have to implement.
type Runnable interface {
	// ID identifies a deployed Runnable.
	ID() string

	// Run tells a deployed Runnable to start working with
	// the given Environment.
	Run(env *Environment) error

	// Stop tells a working Runnable to stop working.
	Stop() error
}
