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
func (c *callback) NotifyRunnerError(id string, rerr error) {
	c.manager.notifyRunnerError(id, rerr)
}

// NotifyRunnerPanic is a callback for the Runner in case of a panic.
func (c *callback) NotifyRunnerPanic(id string, rerr error) {
	c.manager.notifyRunnerPanic(id, rerr)
}

// Logf provides logging for the Manager and via environment to
// the Runnables.
func (c *callback) Logf(format string, v ...interface{}) {
	c.manager.logf(format, v...)
}
