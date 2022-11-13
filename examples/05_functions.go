package main

import "fmt"

// A function that takes two integers and returns their sum as an integer.
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit
// the type name for the like-typed parameters.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Call a function as with other languages using functionName(args).
func main() {
	one, two, three := 1, 2, 3
	result := plus(one, two)
	fmt.Println(one, "+", two, "=", result)

	result = plusPlus(1, 2, 3)
	fmt.Println(one, "+", two, "+", three, "=", result)
}
