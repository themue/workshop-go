// Go Workshop - Practise - Data, Objects, and Services

package api

import (
	"context"

	"github.com/themue/workshop-go/practise/dos/manager"
)

// API is a service providing an external API.
type API struct {
	ctx     context.Context
	manager *manager.Manager
}

// New creates an API service running in the background.
func New(ctx context.Context, manager *manager.Manager) *API {
	return &API{
		ctx:     ctx,
		manager: manager,
	}
}
