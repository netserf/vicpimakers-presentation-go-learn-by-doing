package main

import "fmt"

const keySize = "%-10s"
const fmtSlice = keySize + " %v\n"
const fmtString = keySize + " %s\n"
const fmtInteger = keySize + " %d\n"

func main() {
	// Slices are more commonly used in go rather than arrays.
	// Use the builtin make() function to compose a slice.
	mySlice := make([]string, 3)
	fmt.Printf(fmtSlice, "empty:", mySlice)

	// Set and get slice elements
	mySlice[0] = "a"
	mySlice[1] = "b"
	mySlice[2] = "c"
	fmt.Printf(fmtSlice, "get all:", mySlice)
	fmt.Printf(fmtString, "get [2]:", mySlice[2])

	// len() function available for slices
	fmt.Printf(fmtInteger, "length:", len(mySlice))

	// You can also append() to slices
	mySlice = append(mySlice, "d")
	mySlice = append(mySlice, "e", "f")
	fmt.Printf(fmtSlice, "append:", mySlice)

	// You can make a copy of a slice using copy()
	copySlice := make([]string, len(mySlice))
	copy(copySlice, mySlice)
	fmt.Printf(fmtSlice, "copied:", copySlice)

	// Slices support a slice operator using slice[low:high].
	partialSlice := mySlice[2:5]
	fmt.Printf(fmtSlice, "sliced 1:", partialSlice)

	partialSlice = mySlice[:5]
	fmt.Printf(fmtSlice, "sliced 2:", partialSlice)

	partialSlice = mySlice[2:]
	fmt.Printf(fmtSlice, "sliced 3:", partialSlice)

	// You can also use the blank [] to declare a slice
	shorthandSlice := []string{"g", "h", "i"}
	fmt.Printf(fmtSlice, "shorthand:", shorthandSlice)
}
