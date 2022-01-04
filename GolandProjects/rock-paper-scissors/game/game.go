package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (game *Game) Rounds() {
	// use select to process input the channels
	// print to screen
	// keep track os round number
	for {
		select {
		case round := <-game.RoundChan:
			game.Round.RoundNumber = game.Round.RoundNumber + round
			game.RoundChan <- 1
		case msg := <-game.DisplayChan:
			fmt.Println(msg)
		}
	}
}

// ClearScreen clears the screen
func (game *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (game *Game) PrintIntro() {
	// *** print out some instructions
	fmt.Print("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good Luck!")
	fmt.Println()

}

func (game *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	fmt.Println()
	fmt.Println("Round", game.Round.RoundNumber)
	fmt.Println("------------------")
	fmt.Print("Please enter rock, paper, or scissors -->")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	// *** print out player choice
	fmt.Println()
	game.DisplayChan <- fmt.Sprintf("Player choose %s", strings.ToUpper(playerChoice))

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK!")
		break
	case PAPER:
		fmt.Println("Computer chose PAPER!")
		break
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS!")
		break
	default:
	}

	//fmt.Println("Player value is", playerValue)
	fmt.Println()
	if playerValue == computerValue {
		game.DisplayChan <- "It's a draw!"
		return false
	} else {
		// *** refactor the logic to keep track of score and print messages to two new functions, computerWins and playerWins

		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				game.computerWins()
			} else {
				game.playerWins()
			}
			break
		default:
			// *** decrement i by 1, since we're repeating the round
			game.DisplayChan <- "Invalid choice!"
			return false
		}
	}
	return true
}
func (game *Game) computerWins() {
	game.Round.ComputerScore++
	game.DisplayChan <- "Computer wins!"
}

func (game *Game) playerWins() {
	game.Round.ComputerScore++
	game.DisplayChan <- "Player wins!"
}
