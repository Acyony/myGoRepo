package main

import (
	"log"
)

// Basic Types - numbers, strings, booleans
var myInt int

//var myInt16 int16
//var myInt32 int32
//var myInt64 int64
var myUint uint
var myFloat float32
var myFloat64 float64

// Reference types(pointers, slices, maps, functions, channels)
// Interface type

func main() {
	// Basic Types - numbers
	myInt = 14
	myUint = 35

	myFloat = 11.7
	myFloat64 = 15.9

	log.Println(myInt, myUint, myFloat, myFloat64)
	// Basic Types - strings
	myString := "Alcione"
	log.Println(myString)
	// strings in go are immutable
	myString = "New Name"

	// Basic Types - booleans
	var myBoll = false
	myBoll = true
	log.Println(myBoll)
}
