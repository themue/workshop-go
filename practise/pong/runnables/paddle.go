// Go Workshop - Practise - Pong

package runnables

import (
	"strings"

	"github.com/themue/workshop-go/practise/gube/pkg/runnable"
	"github.com/themue/workshop-go/practise/gube/pkg/storage"
)

// Paddle is an example for a Runnable. It will play
// pong via Storage with another Runnable.
type Paddle struct {
	id        string
	env       *runnable.Environment
	storage   storage.Storage
	opponents []string
}

// ID implements the Runnable interface.
func (p *Paddle) ID() string {
	return p.id
}

// Run implements the Runnable interface.
func (p *Paddle) Run(env *runnable.Environment) error {
	p.env = env
	p.storage = env.Storages[0]
	cfg := env.Configurations[0]
	p.opponents = strings.Split(cfg.Values["targets"], ",")

	p.storage.AddListener(p.id, p.listen)

	return nil
}

// Stop implements the Runnable interface.
func (p *Paddle) Stop() error {
	return nil
}

// listen reacts to a storage event.
func (p *Paddle) listen(op, key, value string) {
	if op != storage.PutOp || key != p.id {
		return
	}
	for _, opponent := range p.opponents {
		p.storage.Put(opponent, value)
	}
}
