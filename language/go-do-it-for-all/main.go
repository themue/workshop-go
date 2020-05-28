// Go Workshop - Language - Goroutines - Go do it for all

package main

import (
	"fmt"
	"sync"
)

// DoItForAll has no output channel but an input channel
// to work with a stream of input values.
func DoItForAll(id int, inc <-chan int, wg *sync.WaitGroup) {
	fmt.Println("----- Start Goroutine DoItForAll", id)
	sum := 0
	for in := range inc {
		sum += in
	}
	fmt.Println("Goroutine DoItForAll", id, "Sum", sum)
	fmt.Println("----- End Goroutine DoItForAll", id)

	// Tell WaitGroup we're done.
	wg.Done()
}

// DoItAsPipe has both, input and output.
func DoItAsPipe(id int, inc <-chan int, outc chan<- int) {
	fmt.Println("----- Start Goroutine DoItAsPipe", id)
	// Loop terminates when input channel gets closed.
	for in := range inc {
		outc <- in * 2
	}
	// Closing output channel to signal that it's done.
	close(outc)
	fmt.Println("----- End Goroutine DoItAsPipe", id)
}

func main() {
	// Create the channels.
	inc := make(chan int)
	outc := make(chan int)

	// WaitGroups are a tool to sync goroutines. Here
	// we add 1 to tell we're waiting for one goroutine.
	var wg sync.WaitGroup

	wg.Add(1)

	// Start the goroutines with the channels.
	go DoItAsPipe(1, inc, outc)
	go DoItForAll(2, outc, &wg)

	// Feed the pipe and then signal the end.
	for i := 1; i <= 5; i++ {
		inc <- i
	}
	close(inc)

	// And now wait.
	wg.Wait()
}
