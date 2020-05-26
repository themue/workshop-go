// Go Workshop - Language - Defer Statements

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ExecutionOrder shows how the deferred functions are
// stacked and called top-down.
func ExecutionOrder() {
	fmt.Println("----- Defer Execution Order")
	fmt.Println("execution order defer: before deferring")

	defer fmt.Println("execution order defer: deferred first, executed last")
	defer fmt.Println("execution order defer: deferred second")
	defer fmt.Println("execution order defer: deferred third")

	fmt.Println("execution order defer: have a break deferring")

	defer fmt.Println("execution order defer: deferred fourth")
	defer fmt.Println("execution order defer: deferred fifth")
	defer fmt.Println("execution order defer: deferred last, executed first")

	fmt.Println("execution order defer: happily done, now let's see")
}

// createTempFile just creates a temporary file for demo.
func createTempFile() (string, error) {
	content := []byte("content of the tempfile")
	tmpfile, err := ioutil.TempFile("", "source")
	if err != nil {
		return "", err
	}
	if _, err := tmpfile.Write(content); err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}

// TypicalUsage shows how defers are typically used.
func TypicalUsage() {
	fmt.Println("----- Defer Typical Usage")
	tmpfileName, err := createTempFile()
	if err != nil {
		fmt.Println("error:", err)
	}
	// First defer for cleanup of temporary file.
	defer func() {
		fmt.Println("- Removing Source")
		os.Remove(tmpfileName)
	}()

	// Now copy from source to destination.
	fmt.Println("- Opening Source:", tmpfileName)
	source, err := os.Open(tmpfileName)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Second defer for closing of temporary file.
	defer func() {
		fmt.Println("- Closing Source")
		source.Close()
	}()

	destination, err := ioutil.TempFile("", "destination")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("- Created Destination:", destination.Name())
	// Third defer for cleanup of temporary destination.
	defer func() {
		fmt.Println("- Removing Destination")
		os.Remove(destination.Name())
	}()
	// Fourth defer for closing of destination.
	defer func() {
		fmt.Println("- Closing Destination")
		destination.Close()
	}()

	n, err := io.Copy(destination, source)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("- Copied Files:", n, "Bytes")
}

func main() {
	ExecutionOrder()
	TypicalUsage()
}
