package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records.
// The &i syntax gives the memory address of i, i.e. a pointer to i.
// Pointers can be printed too.

const keySize = "%-35s"
const fmtInteger = keySize + " %d\n"
const fmtHex = keySize + " %#x\n"

func zeroValue(integerValue int) {
	integerValue = 0
}

func zeroPointer(integerPointer *int) {
	*integerPointer = 0
}

func main() {
	initialInteger := 1
	fmt.Printf(fmtInteger, "initialInteger:", initialInteger)

	// zeroValue has an int parameter and the arguments will be passed by value.
	// So zeroValue gets a copy of intialInteger which ensures initialInteger's
	// original value remains unchanged.
	zeroValue(initialInteger)
	fmt.Printf(fmtInteger, "initialInteger after zeroValue():", initialInteger)

	// &initialInteger gives the memory address of the initialInteger.
	fmt.Printf(fmtHex, "memory address of the pointer:", &initialInteger)

	// zeroPointer has an *int parameter, meaning that it takes an int pointer.
	// The *integerPointer code in the function body dereferences the pointer
	// from its memory address to the current value at that address. So,
	// assigning a value to the dereferenced pointer changes the value at the
	// referenced address.
	zeroPointer(&initialInteger)
	fmt.Printf(fmtInteger, "initialInteger after zeroPointer():",
		initialInteger)
}
