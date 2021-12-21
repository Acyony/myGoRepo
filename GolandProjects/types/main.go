package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types(array, struct)
// reference types(pointers, slices, maps, functions, channels)
// interface type

func main() {
	myInt = 10
	myUint = 20
	myFloat = 10.5
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Alcione"
	log.Println(myString)
	// strings in go are immutables
	myString = "new name"

	var myBool = true
	myBool = false
	log.Println(myBool)
}
