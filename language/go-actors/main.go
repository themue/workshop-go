// Go Workshop - Language - Goroutines - Go Actor

package main

import (
	"context"
	"fmt"
	"sync"

	"tideland.dev/go/audit/generators"

	"github.com/themue/workshop-go/language/go-actors/kvdb"
)

// Perform adds random data to the database.
func Perform(db *kvdb.KVDB, gen *generators.Generator, wg *sync.WaitGroup) {
	defer wg.Done()
	count := gen.Int(50, 100)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("%03d", gen.Int(1, 999))
		value := gen.Word()
		if current := db.Put(key, value); current != nil {
			fmt.Println("Key:", key, "/ Updated:", current, "To:", value)
		}
	}
}

func main() {
	fmt.Println("----- Start")

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
	fmt.Println("----- Begin Listing")
	db.Do(func(key string, value interface{}) {
		fmt.Println("Key:", key, "/ Value:", value)
	})
	value := db.Get("555")
	fmt.Println("Key: 555 / Value:", value)
	fmt.Println("----- End Listing")

	// Terminate database.
	cancel()

	fmt.Println("----- Done")
}
