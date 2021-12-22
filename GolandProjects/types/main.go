package main

import "fmt"

// reference types(pointers, slices, maps, functions, channels)
// interface type

func main() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//delete(intMap, "four")

	// finding if an element exist in a map
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	// change the value of an element in a map
	intMap["two"] = 4
}

// not possible to sort a map
