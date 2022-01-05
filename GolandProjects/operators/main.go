package main

import (
	"fmt"
	"math"
)

func main() {
	/*answer := 7 + 3*8
	fmt.Println("Answer =", answer)
	answer = (7 + 3) * 8
	fmt.Println("Answer is now =", answer)*/

	// multiplication
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area is =", area)

	// integer division
	half := 1 / 2
	fmt.Println("Half integer division =", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("Half float is =", halfFloat)

	// squaring (raising to the power)
	badThreeSquared := 3 ^ 2
	fmt.Println("Bad Three Squared =", badThreeSquared)
	fmt.Println("-------=^.^=-------")
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Good Three Squared =", goodThreeSquared)

	// modules
	remainder := 50 % 3
	fmt.Println("remainder =", remainder)

	// unary operators

	x := 3
	x++
	fmt.Println("x is now =", x)

	x--
	x--
	fmt.Println("x is now =", x)
}
