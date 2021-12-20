package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatalln(err)
	}

	// defer keyword will be executed only after the current function ends his execution
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocca"
	coffees[5] = "Macchiatto"
	coffees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("----=^.^=-----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocca")
	fmt.Println("5 - Macchiatto")
	fmt.Println("6 - Expresso")
	fmt.Println("7 - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		integer, _ := strconv.Atoi(string(char))

		//%q is a placeholder for room
		//%d is a placeholder for integer
		//%s is a placeholder for string
		fmt.Println(fmt.Sprintf("You chose %s", coffees[integer]))

		if char == 'q' || char == 'Q' {
			break
		}
	}

	fmt.Println("Program exiting...")
}
