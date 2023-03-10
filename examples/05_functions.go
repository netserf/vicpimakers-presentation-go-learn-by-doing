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

// Functions can also return multiple values.
func plusAndMinus(a, b int) (int, int) {
	return a + b, a - b
}

// Function can take variadic parameters - i.e. an arbitrary number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Call a function as with other languages using functionName(args).
func main() {
	fmt.Println("1 + 2 =", plus(1, 2))

	fmt.Println("1 + 2 + 3 =", plusPlus(1, 2, 3))

	plusResult, minusResult := plusAndMinus(5, 2)
	fmt.Println("5 + 2 =", plusResult)
	fmt.Println("5 - 2 =", minusResult)

	fmt.Println("1 + 2 + 3 + 4 + 5 =", sum(1, 2, 3, 4, 5))
}
