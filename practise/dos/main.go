// Go Workshop - Practise - Data, Objects, and Services

package main

import (
	"context"

	"github.com/themue/workshop-go/practise/dos/api"
	"github.com/themue/workshop-go/practise/dos/manager"
	"github.com/themue/workshop-go/practise/dos/registry"
)

func main() {
	ctx := context.Background()
	manager := manager.New(ctx, registry.NewSimpleRegistry(), nil)
	_ = api.New(ctx, manager)

	<-ctx.Done()
}
