package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

//rune = data type used to build string

var keyPressChan chan rune

func main() {
	// channels allow the program to send data from one place to another. Used go routine
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
