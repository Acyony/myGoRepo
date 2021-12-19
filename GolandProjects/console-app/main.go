package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
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

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program exiting...")
}
