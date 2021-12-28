package main

import "fmt"

// Expression is a value that can be evaluated to a single value

func main() {
	age := 24
	name := "Jack"
	rightHanded := true
	fmt.Printf("%s is %d years old. RightHanded %t", name, age, rightHanded)
	fmt.Println("---------=^.^=-----------")

	//ageInTenYears is an expression that can be evaluated to a single value
	ageInTenYears := age + 10
	fmt.Printf("In ten years %s will be %d years old", name, ageInTenYears)
	fmt.Println("---------=^.^=-----------")

	isATeenAge := age >= 13
	fmt.Println(name, "is a teenager", isATeenAge)
	fmt.Println("---------=^.^=-----------")

}
