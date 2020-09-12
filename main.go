// Tic-tac-toe game in go.

package main

import (
	"fmt"
)

func main() {
	board := Board{}

	choices := [2]string{"O", "X"}
	cpuSymbol := choices[GenerateRandomNumber(0, 2)]
	board.Initialise(cpuSymbol)

	for {
		board.PrintBoard()

		move := ReadMove()
		moveErr := board.Move(move)

		if moveErr != nil {
			fmt.Println("Invalid move")
		}

		hasWinner, winner := board.GetWinner()

		if hasWinner && winner == " " {
			board.PrintBoard()
			fmt.Println("Draw!")
			break
		}

		if hasWinner {
			board.PrintBoard()
			fmt.Println(winner + " Wins!")
			break
		}
	}
}
