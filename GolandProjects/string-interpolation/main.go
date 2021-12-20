package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
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
	OwnsADog  bool
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
	user.OwnsADog = readBoll("Do you own a dog (y/n)?")
	// This way uses less memory, and it's much faster
	fmt.Printf("Your name is %s, and you are %d years old. Your favourite number is %.2f. Owns a dog: %t.",
		user.UserName,
		user.Age,
		user.FavNumber,
		user.OwnsADog,
	)
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

// to read a boolean

func readBoll(str string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatalln(err)
	}

	// to close the keyboard

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(str)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type  y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}

}
