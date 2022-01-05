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
			game.DisplayChan <- ""
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
	game.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
----------------------
Game is played for three rounds, and best of three wins the game. Good Luck!
`)
	<-game.DisplayChan
}

func (game *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	game.DisplayChan <- fmt.Sprintf(`
Round %d
------------------
`, game.Round.RoundNumber)
	<-game.DisplayChan

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
	game.DisplayChan <- ""
	<-game.DisplayChan

	switch computerValue {
	case ROCK:
		game.DisplayChan <- "Computer chose ROCK!"
		<-game.DisplayChan
		break
	case PAPER:
		game.DisplayChan <- "Computer chose PAPER!"
		<-game.DisplayChan
		break
	case SCISSORS:
		game.DisplayChan <- "Computer chose SCISSORS!"
		<-game.DisplayChan
		break
	default:
	}

	//fmt.Println("Player value is", playerValue)
	game.DisplayChan <- ""
	<-game.DisplayChan

	if playerValue == computerValue {
		game.DisplayChan <- "It's a draw!"
		<-game.DisplayChan
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
			<-game.DisplayChan
			return false
		}
	}
	return true
}
func (game *Game) computerWins() {
	game.Round.ComputerScore++
	game.DisplayChan <- "Computer wins!"
	<-game.DisplayChan

}

func (game *Game) playerWins() {
	game.Round.ComputerScore++
	game.DisplayChan <- "Player wins!"
	<-game.DisplayChan

}

func (game *Game) PrintSummary() {

	fmt.Println("Final score")
	fmt.Println("--------------------")
	fmt.Printf("Player: %d/3, Computer %d/3", game.Round.PlayerScore, game.Round.ComputerScore)
	fmt.Println()
	if game.Round.PlayerScore > game.Round.ComputerScore {
		game.DisplayChan <- "Player wins the game!"
		<-game.DisplayChan
	} else {
		game.DisplayChan <- "Computer wins the game!"
		<-game.DisplayChan
	}
}
