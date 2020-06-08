// Go Workshop - Practise - Gube - Manager

package manager

// callback is an invisible type providing public methods
// Runner and Environment expect and which need access to
// the manager. The type and its methods are invisible in
// Go docs.
type callback struct {
	manager *Manager
}

// NotifyRunnerError is a callback for the runner in case of an error.
func (c *callback) NotifyRunnerError(id string, err error) {
	// TODO Restart if configured.
}

// NotifyRunnerPanic is a callback for the Runner in case of a panic.
func (c *callback) NotifyRunnerPanic(id string, err error) {
	c.Logf("service %q had a panic: %v", id, err)
	esvc, err := c.manager.registry.RetrieveService(id)
	if err != nil {
		c.Logf("panic handling: service %q cannot be retrieved: %v", id, err)
		return
	}
	if !esvc.Restart {
		return
	}
	// Restart is configured, so do it.
	err = m.Spawn(id)
	if err != nil {
		c.Logf("panic handling: service %q cannot be restarted: %v", id, err)
	}
}

// Logf provides logging for the Manager and via environment to
// the Runnables.
func (c *callback) Logf(format string, v ...interface{}) {
	c.manager.logger.Printf(format, v...)
}
