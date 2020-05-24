// Go Workshop - Language - Arguments and exit code
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var how string
	var who string

	// Package flag provides flag parsing, default values, and help.
	flag.StringVar(&how, "how", "Hello", "Tell how to greet the greeted.")
	flag.StringVar(&who, "who", "", "Tell who you want to greet.")
	flag.Parse()

	if who == "" {
		fmt.Println("Error: need to know who you want to greet")
		// Package os provides functions to access operating system,
		// function Exit() terminates the program, here with code
		// 1.
		os.Exit(1)
	}

	fmt.Printf("%s, %s\n", how, who)
}
