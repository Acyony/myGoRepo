package main

import (
	"bufio"
	"fmt"
	"myapp/myLogger"
	"os"
	"time"
)

func main() {
	// read input from the user 5 times and write it to a log
	reader := bufio.NewReader(os.Stdin)
	// creating a channel
	ch := make(chan string)

	go myLogger.ListenForLog(ch)

	fmt.Println("Enter something! =^.^= ")

	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
