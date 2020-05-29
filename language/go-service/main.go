// Go Workshop - Language - Goroutines - Go Service

package main

import (
	"context"
	"fmt"
	"sync"

	"tideland.dev/go/audit/generators"

	"github.com/themue/workshop-go/language/go-service/kvdb"
)

// Perform adds random data to the database.
func Perform(db *kvdb.KVDB, gen *generators.Generator, wg *sync.WaitGroup) {
	defer wg.Done()
	count := gen.Int(50, 100)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("%03d", gen.Int(1, 999))
		value := gen.Word()
		db.Put(key, value)
	}
}

func main() {
	// Starting the database with a canceable context.
	gen := generators.New(generators.FixedRand())
	ctx, cancel := context.WithCancel(context.Background())
	db := kvdb.New(ctx)

	// Run a number of goroutines.
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Perform(db, gen, &wg)
	}
	wg.Wait()

	// Check content.
	db.Do(func(key string, value interface{}) {
		fmt.Println("Key:", key, "/ Value:", value)
	})

	value := db.Get("555")
	fmt.Println("Key: 555 / Value:", value)

	// Terminate database.
	cancel()
}
