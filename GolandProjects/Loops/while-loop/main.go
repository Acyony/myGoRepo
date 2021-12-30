package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//to get a random number
	rand.Seed(time.Now().UnixNano())
	// executes a loop while i > 100 - for example
	i := 1000
	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1
		fmt.Println("i is:", i)
		if i > 100 {
			fmt.Println("i is", i, "so the loop keeps going!")
		}
	}
	fmt.Println("Got", i, "and broke out of loop!")
}
