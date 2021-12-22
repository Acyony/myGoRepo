package main

import (
	"fmt"
	"sort"
)

// reference types(pointers, slices, maps, functions, channels)
// interface type

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	fmt.Println(animals)

	// _, x index and the current item in the slice i'm looking at

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")

	// sorting slices strings
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func DeleteFromSlice(a []string, i int) []string {
	//deleting an element using the index
	a[i] = a[len(a)-1] // copy the last element of the slice
	a[len(a)-1] = ""   // given the last element a default value
	a = a[:len(a)-1]   // deleting the last element
	return a
}
