package main

import "fmt"

// Ranges iterate over elements in a variety of data structures.
func main() {

	// Declare an array
	numbersArray := []int{2, 3, 4}
	sum := 0
	// Use a range to iterate over the array of numbers.
	// A range returns both index and value. In this case we don't need the
	// range so we use the blank (_) identifier to indicate we're ignoring it.
	for _, num := range numbersArray {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Here we use both the index and value
	for index, num := range numbersArray {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}

	// Use a range to iterate over a map
	fruitMap := map[string]string{"a": "apple", "b": "banana"}
	for key, fruit := range fruitMap {
		fmt.Printf("%s -> %s\n", key, fruit)
	}

	// A range can also be used to iterate over only the keys in a map
	for k := range fruitMap {
		fmt.Println("key:", k)
	}

	// A range on strings iterates over Unicode code points.
	// The first value is the byte index of the rune and the second the rune
	// itself.
	// TODO - Consider removing from examples. Have not introduced runes yet.
	for index, rune := range "go" {
		fmt.Println(index, rune)
	}
}
