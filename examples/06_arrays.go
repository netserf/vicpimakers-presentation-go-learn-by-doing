package main

import "fmt"

const keySize = "%-10s"
const fmtArray = keySize + " %v\n"
const fmtString = keySize + " %s\n"
const fmtInteger = keySize + " %d\n"

func main() {

	// Create an array that can hold 5 ints.
	var myArray [5]int
	fmt.Printf(fmtArray, "original: ", myArray)

	// Set a value at an index.
	myArray[4] = 100
	fmt.Printf(fmtArray, "set [4]:  ", myArray)
	fmt.Printf(fmtInteger, "get [4]:  ", myArray[4])

	// len() returns the length of an array.
	fmt.Printf(fmtInteger, "length:   ", len(myArray))

	// Declare and initialize an array in one line.
	myArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf(fmtArray, "declared: ", myArray2)

	// Declare and initialize a 2D array.
	var my2dArray [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			my2dArray[i][j] = fmt.Sprintf("(%d, %d) ", i, j)
		}
	}
	// Print the positions of each element.
	fmt.Printf(fmtArray, "2D row 0: ", my2dArray[0])
	fmt.Printf(fmtArray, "2D row 1: ", my2dArray[1])
}
