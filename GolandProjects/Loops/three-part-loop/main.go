package main

import "fmt"

func main() {
	// 1 - part: i := 10 - where the index begins
	// 2 - part: i > 0 - the condition (boolean expression)
	// 2 - part: i-- / i++ - the decrement or increment
	for i := 10; i > 0; i-- {
		fmt.Println("i is:", i)
	}
}
