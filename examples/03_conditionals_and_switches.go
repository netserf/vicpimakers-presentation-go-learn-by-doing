package main

import "fmt"

func main() {
	num := 5
	// You can use if, else if, else statements to execute different code
	// blocks based on conditions.
	fmt.Println("Go supports if, else if, else statements:")
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Go also supports switch / case statements to handle multiple
	// cases for a single variable.
	fmt.Println("Go also supports switch / case statements:")
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown")
	}

	// If you want C-style fall-through behavior, use the fallthrough statement.
	fmt.Println("Go supports an explicit fallthrough statement:")
	s := "two"
	fmt.Println("Show fallthrough for :", s)
	switch {
	case s == "one":
		fmt.Println("1")
		fallthrough
	case s == "two":
		fmt.Println("2")
		fallthrough
	case s == "three":
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("always shown")
	}
}
