package main

import "fmt"

const keySize = "%-7s"
const fmtInteger = keySize + " %d\n"

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
	fmt.Printf(fmtInteger, "sum:", sum)

	// Here we use both the index and value
	for index, num := range numbersArray {
		if num == 3 {
			fmt.Printf(fmtInteger, "index:", index)
		}
	}

	// Use a range to iterate over a map
	fruitMap := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("fruitMap:")
	for key, fruit := range fruitMap {
		fmt.Printf("\t%s -> %s\n", key, fruit)
	}

	// A range can also be used to iterate over only the keys in a map
	fmt.Println("fruitMap keys:")
	for k := range fruitMap {
		fmt.Println("\tkey:", k)
	}
}
