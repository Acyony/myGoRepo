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

// creating a new type of variable => User
type User struct {
	UserName  string
	Age       int
	FavNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	if user.UserName == "" {
		fmt.Println("Please enter a valid name!")
		return
	}

	user.Age = readInt("How old are you?")
	user.FavNumber = readFloat("What is your favourite number?")

	// This way uses less memory, and it's much faster
	fmt.Printf("Your name is %s, and you are %d years old. Your favourite number is %.2f.",
		user.UserName,
		user.Age,
		user.FavNumber)
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

func readFloat(str string) float64 {
	for {
		fmt.Println(str)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		// convert str into a number
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number!")
		} else {
			return num
		}
	}
}
