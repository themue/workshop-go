// Go Workshop - Language - Goroutines - Go do it with me

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// DoItWithMe is a player of a ping-pong game.
func DoItWithMe(id string, inc, outc chan string, wg *sync.WaitGroup) {
	fmt.Println("----- Start Goroutine DoItWithMe", id)

	defer fmt.Println("----- End Goroutine DoItWithMe", id)
	defer wg.Done()
	defer close(outc)

	coin := func() bool {
		i := rand.Intn(100)
		return i > 95
	}

	for in := range inc {
		if coin() {
			// I give up.
			fmt.Println("Goroutine DoItWithMe", id, "gives up")
			return
		}
		//  Return the ball.
		fmt.Println("Goroutine DoItWithMe", id, "plays", in)

		outc <- id
	}
}

func main() {
	// Create the channels.
	pingc := make(chan string)
	pongc := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	// Start the goroutines with the channels.
	go DoItWithMe("ping", pongc, pingc, &wg)
	go DoItWithMe("pong", pingc, pongc, &wg)

	// Service ping.
	pingc <- "ping"

	wg.Wait()
}
