package main

import (
	"fmt"
	"strings"
)

func main() {
	/*fmt.Println()
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
	fmt.Println(sb.String())*/

	//indexes start at 0
	/*
		courseName := "Learning Go for Beginners"
		fmt.Println(courseName[0])
		fmt.Println(string(courseName[7]))

		for i := 0; i <= 24; i++ {
			fmt.Print(string(courseName[i]))
		}
		fmt.Println()
		fmt.Println("-------=^.^=-------")
		var mySlice []string
		mySlice = append(mySlice, "one")
		mySlice = append(mySlice, "two")
		mySlice = append(mySlice, "three")
		mySlice = append(mySlice, "Four")

		fmt.Println("mySlice has", len(mySlice), "elements")
		fmt.Println("The last element is mySlice is:", mySlice[len(mySlice)-1])*/

	/*courses := []string{
		"Learn Go for Beginners crash course",
		"Learn Java for Beginners crash course",
		"Learn Python for Beginners crash course",
		"Learn C for Beginners crash course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in =>", x, "and index is:", strings.Index(x, "Go"))
		}
	}

	newString := "Go is an amazing language to learn and the Go gopher is an iconic mascot!"
	// the method HasPrefix returns a bool
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "gopher"))
	fmt.Println(strings.Count(newString, "cat"))
	fmt.Println(strings.Index(newString, "gopher"))
	fmt.Println(strings.LastIndex(newString, "gopher"))*/

	/*newString := "Go is an amazing language to learn and the Go gopher is an iconic mascot!"
	if strings.Contains(newString, "Go") {
		//newString = strings.Replace(newString, "Go", "Golang", 1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}
	fmt.Println(newString)

	badEmail := " me@mail.com "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Println("=%5=", badEmail)
	fmt.Println()*/

	str := "gopher gopher gopher gopher gopher"
	str = replaceNth(str, "gopher", "Golang", 4)
	fmt.Println(str)
}

func replaceNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// we did not find it
			break
		}

		// have found it
		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
