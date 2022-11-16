package main

import "fmt"

func main() {
	// Slices are more commonly used in go rather than arrays.
	// Use the builtin make() function to compose a slice.
	mySlice := make([]string, 3)
	fmt.Println("empty:    ", mySlice)

	// Set and get slice elements
	mySlice[0] = "a"
	mySlice[1] = "b"
	mySlice[2] = "c"
	fmt.Println("get all:  ", mySlice)
	fmt.Println("get [2]:  ", mySlice[2])

	// len() function available for slices
	fmt.Println("length:   ", len(mySlice))

	// You can also append() to slices
	mySlice = append(mySlice, "d")
	mySlice = append(mySlice, "e", "f")
	fmt.Println("append:   ", mySlice)

	// You can make a copy of a slice using copy()
	copySlice := make([]string, len(mySlice))
	copy(copySlice, mySlice)
	fmt.Println("copied:   ", copySlice)

	// Slices support a slice operator using slice[low:high].
	partialSlice := mySlice[2:5]
	fmt.Println("sliced 1: ", partialSlice)

	partialSlice = mySlice[:5]
	fmt.Println("sliced 2: ", partialSlice)

	partialSlice = mySlice[2:]
	fmt.Println("sliced 3: ", partialSlice)

	// You can also use the blank [] to declare a slice
	shorthandSlice := []string{"g", "h", "i"}
	fmt.Println("shorthand:", shorthandSlice)
}
