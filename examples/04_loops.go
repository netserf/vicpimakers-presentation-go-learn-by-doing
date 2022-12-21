package main

import "fmt"

// We only have for loops in Go, but they are powerful.
// This is a small subset of what you can do with them.

func main() {

	// A basic, single-condition for loop.
	fmt.Print("Loop 1:  ")
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	fmt.Print("\n\nLoop 2:  ")
	for j := 7; j <= 9; j++ {
		fmt.Print(j, " ")
	}

	// A for loop without a condition will loop repeatedly a break statement.
	fmt.Print("\n\nLoop 3:  ")
	for {
		fmt.Println("loop until break")
		break
	}

	// You can also use the continue statement to skip to the next iteration.
	fmt.Print("\nLoop 4:  ")
	for k := 0; k <= 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Print(k, " ")
	}
}
