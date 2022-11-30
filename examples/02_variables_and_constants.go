package main

import "fmt"

const constantString1 string = "my constant"
const constantInteger1 int = 1000

func main() {

	// Go supports constants
	fmt.Println("constantString1:", constantString1)
	fmt.Println("constantInteger1:", constantInteger1)

	// A constant is a simple, immutable value.
	// constantInteger1 = 2000  // Error: cannot assign to constantInteger1

	// Variables are explicitly declared, and can be explicitly typed
	var string1 string = "1st string is explicitly typed"
	fmt.Println("string1:", string1)

	// OR implicitly typed
	var string2 = "2nd string is implicitly typed"
	fmt.Println("string2:", string2)

	// Variable can also be declared with the := syntax
	string3 := "3rd string is declared with the := shorthand"
	fmt.Println("string3:", string3)

	// Multiple variables can be declared on the same line
	var int1, int2 int = 1, 2
	fmt.Println("integers:", int1, int2)

	// Booleans are supported
	var bool1 = true
	fmt.Println("boolean:", bool1)

	// Floats are supported
	var float1 = 1.2
	fmt.Println("float:", float1)
}
