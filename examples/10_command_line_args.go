package main

import (
	"fmt"
	"os"
)

// You can get individual args with normal indexing.
// To experiment with command-line arguments it’s best to build a binary with go build first.

func main() {
	// os.Args provides access to command-line arguments as a slice.

	// The first value in the slice is the path to the program.
	executableName := os.Args[0]
	fmt.Println("executable:", executableName)

	// os.Args[1:] holds only the arguments to the program.
	arguments := os.Args[1:]
	fmt.Println("arguments:", arguments)

	// fetch a single argument
	if len(arguments) > 0 {
		fmt.Println("1st arg:", arguments[0])
	}
}
