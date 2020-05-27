// Go Workshop - Language - Goroutines - Return the Result

package main

import (
	"fmt"
)

// DoIt here also takes a channel to return a result
// to the caller. Another common pattern.
func DoIt(id, count int, resultc chan int) {
	fmt.Println("----- Start Goroutine", id)
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	resultc <- sum
	fmt.Println("----- End Goroutine", id)
}

func main() {
	resultc1 := make(chan int)
	resultc2 := make(chan int)

	go DoIt(1, 25, resultc1)
	go DoIt(2, 50, resultc2)

	sum1 := <-resultc1
	sum2 := <-resultc2

	fmt.Println("Result 1", sum1)
	fmt.Println("Result 2", sum2)
}
