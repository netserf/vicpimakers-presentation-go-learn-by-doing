package main

import "fmt"

const constantString1 string = "my constant"
const constantInteger1 int = 1000

func main() {
	// variables are explicitly declared, and they can be typed explicitly
	var string1 string = "1st string is explicitly typed"
	fmt.Println("string1:", string1)
	// ... OR implicitly
	var string2 = "2nd string is implicitly typed"
	fmt.Println("string2:", string2)

	// shorthand declaration
	string3 := "3rd string is declared with the := shorthand"
	fmt.Println("string3:", string3)

	// multiple variables can be declared on the same line
	var int1, int2 int = 1, 2
	fmt.Println("integers:", int1, int2)

	// booleans
	var bool1 = true
	fmt.Println("boolean:", bool1)

	// floats
	var float1 = 1.2
	fmt.Println("float:", float1)

	// constants
	fmt.Println("constantString1:", constantString1)
	fmt.Println("constantInteger1:", constantInteger1)
}
