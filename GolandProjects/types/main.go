package main

import "fmt"

// reference types(pointers, slices, maps, functions, channels)
// pointers => something that points to a specific location  in memory
// interface type

func main() {
	x := 10
	// hold the address where x is storage in the memory
	// &x reference to a pointer
	// *myFirstPointer to get the value and change it

	myFirstPointer := &x
	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)
	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
