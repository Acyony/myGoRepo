package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// package level var
var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")
	if userName == "" {
		fmt.Println("Please enter a valid name!")
		return
	}

	age := readInt("How old are you?")
	fmt.Println("Your name is: ", userName, ". You are", age, "years old.")
}

func prompt() {
	fmt.Println("=> ")
}

// to read userName
func readString(str string) string {
	for {
		fmt.Println(str)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		// convert number into a string
		_, err := strconv.Atoi(userInput)
		if err == nil {
			return ""
		}

		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

// to read userAge

func readInt(str string) int {
	for {
		fmt.Println(str)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		// convert str into a number
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number!")
		} else {
			return num
		}
	}
}
