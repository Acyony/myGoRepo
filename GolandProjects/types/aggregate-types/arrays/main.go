package main

import "log"

// Aggregate Types - arrays

func main() {
	// arrays
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"
	log.Println("First element in the array is:", myStrings[0])
}
