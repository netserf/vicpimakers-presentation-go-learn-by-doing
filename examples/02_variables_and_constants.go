package main

import "fmt"

const constString1 string = "my constant" // Can declare the type
const constInteger1 int = 1000
const keySize = "%-15s" // Can also leave the type out
const fmtString = keySize + " %s\n"
const fmtInteger = keySize + " %d\n"
const fmt2Integers = keySize + " %d %d\n"
const fmtBool = keySize + " %t\n"
const fmtFloat = keySize + " %g\n"

func main() {
	// Go supports constants
	fmt.Printf(fmtString, "constString1:", constString1)
	fmt.Printf(fmtInteger, "constInteger1:", constInteger1)

	// A constant is a simple, immutable value.
	// constantInteger1 = 2000  // Error: cannot assign to constantInteger1

	// Variables are explicitly declared, and can be explicitly typed
	var string1 string = "1st string is explicitly typed"
	fmt.Printf(fmtString, "string1:", string1)

	// OR implicitly typed
	var string2 = "2nd string is implicitly typed"
	fmt.Printf(fmtString, "string2:", string2)

	// Variable can also be declared with the := syntax
	string3 := "3rd string is declared with the := shorthand"
	fmt.Printf(fmtString, "string3:", string3)

	// Multiple variables can be declared on the same line
	var int1, int2 int = 1, 2
	fmt.Printf(fmt2Integers, "integers:", int1, int2)

	// Booleans are supported
	var bool1 = true
	fmt.Printf(fmtBool, "boolean:", bool1)

	// Floats are supported
	var float1 = 1.2
	fmt.Printf(fmtFloat, "float:", float1)
}
