package main

import "fmt"

func main() {
	apples := 25
	oranges := 10
	//boolean expression
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > < >= <=
	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()
	if apples > oranges {
		fmt.Println("You have more apples than oranges!")
	} else {
		fmt.Println("You have more oranges than apples!")
	}
}
