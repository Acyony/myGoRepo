package main

import "fmt"

// aggregate types(array, struct)
// reference types(pointers, slices, maps, functions, channels)
// interface type

func main() {
	// declaring array in goLang
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("The first element in thr array is:", myStrings[0])
}
