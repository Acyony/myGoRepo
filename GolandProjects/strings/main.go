package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world!"
	fmt.Println("String", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x  ", name[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("Three ways to concatenate strings")
	h := "hello, "
	w := "world."

	// using + not very efficient
	myString := h + " " + w
	fmt.Println(myString)

	// using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using string builder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())
}
