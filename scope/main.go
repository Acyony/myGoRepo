package main

import (
	"fmt"
	"myapp/packageone"
)

// package level variable - global scope
var one = "One"

func main() {
	var oneLocalVar = "This is a block level variable"
	fmt.Println(oneLocalVar)

	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From package one", newString)
}

func myFunc() {
	fmt.Println(one)
}
