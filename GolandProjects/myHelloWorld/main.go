// main => executable file conventional called main
package main

import (
	"bufio"
	"fmt"
	"myHelloWorld/doctor"
	"os"
	"strings"
)

//fmt stands for format

// entry point of the executable application
func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	// for loop
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		// Windows
		// First argument userInput -> What the user types
		// Second argument is what I'm looking for "\r\n"
		// Third argument is what I want to replace with ""
		// Fourth argument is how many times I want to it -1. That means replace everywhere it will be founded

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		// Other operator systems
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			// type break to finish the loop
			break
		} else {
			// print the response
			fmt.Println(doctor.Response(userInput))
		}
	}
}
