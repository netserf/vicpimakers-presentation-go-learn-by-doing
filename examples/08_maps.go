package main

import "fmt"

func main() {
	// To create an empty map, use the builtin make() function
	myMap := make(map[string]int)

	// Set key/value pairs with the name[key] = value syntax
	myMap["key1"] = 7
	myMap["key2"] = 13

	// Output the entire map
	fmt.Println("map:    ", myMap)

	// Output a single value from the map
	fmt.Println("value:  ", myMap["key1"])

	// Discover the map length
	fmt.Println("length: ", len(myMap))

	// Use the delete() function to remove an element from the map
	delete(myMap, "key2")
	fmt.Println("map:    ", myMap)

	// There is a 2nd returned value in the map to indicate if it was present
	_, present := myMap["key2"]
	fmt.Println("present:", present)

	// shorthand map declaration
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("short:  ", n)
}
