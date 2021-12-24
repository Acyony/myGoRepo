package main

import "fmt"

func main() {
	// points to a specific location in the memory
	x := 10

	myFirstPointer := &x // it's point to the location where 10 is located
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
