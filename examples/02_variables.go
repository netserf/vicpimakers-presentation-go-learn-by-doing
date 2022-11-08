package main

import "fmt"

func main() {
	// variables are explicitly declare, but can be typed implicitly or
	// explicitly typed
	var my_string_1_explicit string = "1st string - explicit"
	fmt.Println(my_string_1_explicit)
	var my_string_2_implicit = "2nd string - implicit"
	fmt.Println(my_string_2_implicit)

	// shorthand declaration
	my_string_3 := "3rd string - shorthand"
	fmt.Println(my_string_3)

	// multiple variables can be declared on the same line
	var my_int_1, my_int_2 int = 1, 2
	fmt.Println(my_int_1, my_int_2)

	// booleans
	var my_bool_1 = true
	fmt.Println(my_bool_1)

	// floats
	var my_float_1 = 1.2
	fmt.Println(my_float_1)
}
