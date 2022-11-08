package main

import "fmt"

const constant_string_1 string = "my constant"
const constant_integer_1 int = 1000

func main() {
	// variables are explicitly declared, and they can be typed explicitly
	var string_1 string = "1st string is explicitly typed"
	fmt.Println("string_1:", string_1)
	// ... OR implicitly
	var string_2 = "2nd string is implicitly typed"
	fmt.Println("string_2:", string_2)

	// shorthand declaration
	string_3 := "3rd string is declared with the := shorthand"
	fmt.Println("string_3:", string_3)

	// multiple variables can be declared on the same line
	var int_1, int_2 int = 1, 2
	fmt.Println("integers:", int_1, int_2)

	// booleans
	var bool_1 = true
	fmt.Println("boolean:", bool_1)

	// floats
	var float_1 = 1.2
	fmt.Println("float:", float_1)

	// constants
	fmt.Println("constant_string_1:", constant_string_1)
	fmt.Println("constant_integer_1:", constant_integer_1)
}
