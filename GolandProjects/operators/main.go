package main

func main() {
	/*answer := 7 + 3*8
	fmt.Println("Answer =", answer)
	answer = (7 + 3) * 8
	fmt.Println("Answer is now =", answer)*/
	/*
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
		fmt.Println("x is now =", x)*/

	// reference source
	// https://www.tutorialspoint.com/go/go_operators_precedence.htm
	// https://www.golangprograms.com/arithmetic-operators-in-go-programming-language.html
	// https://go.dev/ref/spec#Arithmetic_operators

	/*// precedents

	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)
	fmt.Println("a =", a, "b =", b, "c =", c)

	// integer division
	unclear := 12 * (3 / 4)
	fmt.Println("unclear =", unclear)

	// parenthesis
	foo := 12.0 / 3.0 / 40
	fmt.Println("foo =", foo)
	fmt.Println("-------=^.^=-------")

	foo1 := 12.0 / (3.0 / 40)
	fmt.Println("foo1 =", foo1)

	// addition/subtraction
	fmt.Println()

	x := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("x =", x, "y =", y, "z =", z)
	fmt.Println("-------=^.^=-------")

	x = 12 + 3*4
	y = (12 + 3) * 4
	z = 12 + (3 * 4)
	fmt.Println("x =", x, "y =", y, "z =", z)*/

	/*// modulus operator

	x := 12
	y := 5
	if x%y == 0 {
		fmt.Println(y, "divide exactly into ", x)
	} else {
		fmt.Println(y, " doesn't divide exactly into ", x)
	}

	thisMonth := 1
	fmt.Println("The month after", thisMonth, "is", thisMonth+1)

	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12+1)
	}*/

	// relational and conditional operators

	/*second := 31
	minute := 1

	if (minute < 59) && ((second + 1) > 59) {
		minute++
	}

	fmt.Println(minute)*/

	/*a := 12
	b := 5*/
	/*
		if b != 0 {
			c := divideTwoNumbers(a, b)

			if c == 2 {
				fmt.Println("we found a two")
			}
		} else {
			fmt.Println("You are dividing by 0!")
		}*/

	/*c, err := divideTwoNumbers(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2!")
		}
	}*/

}

/*
func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}*/
