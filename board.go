package main

import (
	"errors"
	"fmt"
)

// Board represents the game board
type Board struct {
	PlayerSymbol string
	CPUSymbol    string
	State        [9]string
}

// Initialise performs board setup
func (board *Board) Initialise(cpuSymbol string) {
	for i := 0; i < 9; i++ {
		board.State[i] = " "
	}

	if cpuSymbol == "X" {
		board.CPUSymbol = "X"
		board.PlayerSymbol = "O"
		move := GenerateMove(*board)
		board.State[move] = "X"
	} else {
		board.CPUSymbol = "O"
		board.PlayerSymbol = "X"
	}
}

// PrintBoard displays board state
func (board *Board) PrintBoard() {
	ClearScreen()
	fmt.Println("")

	fmt.Println("Player: ", board.PlayerSymbol)
	fmt.Println("CPU: ", board.CPUSymbol)
	fmt.Println("")

	fmt.Println(" --- --- --- ")
	fmt.Println("| " + board.State[0] + " | " + board.State[1] + " | " + board.State[2] + " |")
	fmt.Println(" --- --- --- ")
	fmt.Println("| " + board.State[3] + " | " + board.State[4] + " | " + board.State[5] + " |")
	fmt.Println(" --- --- --- ")
	fmt.Println("| " + board.State[6] + " | " + board.State[7] + " | " + board.State[8] + " |")
	fmt.Println(" --- --- --- ")

	fmt.Println("")
}

// IsValidMove Checks if the move is valid
func (board *Board) IsValidMove(move int) bool {
	if move > 8 || move < 0 {
		return false
	}

	return board.State[move] == " "
}

// GetWinner searches for a winner
func (board Board) GetWinner() (bool, string) {
	lines := [8][3]int{
		[3]int{0, 1, 2},
		[3]int{3, 4, 5},
		[3]int{6, 7, 8},
		[3]int{0, 3, 6},
		[3]int{1, 4, 7},
		[3]int{2, 5, 8},
		[3]int{0, 4, 8},
		[3]int{2, 4, 6},
	}

	// Check winner
	for _, line := range lines {
		if IsFullLine(board.State[line[0]], board.State[line[1]], board.State[line[2]]) {
			return true, board.State[line[0]]
		}
	}

	// Check draw
	count := 0
	for _, symbol := range board.State {
		if symbol != " " {
			count++
		}
	}

	if count == 9 {
		return true, " "
	}

	// No winner
	return false, ""
}

// Move logs player move and calculate next CPU move
func (board *Board) Move(move int) error {
	if !board.IsValidMove(move) {
		return errors.New("Invalid move")
	}

	board.State[move] = board.PlayerSymbol

	// Do not calculate next move there is a winner already.
	if hasWinner, _ := board.GetWinner(); hasWinner {
		return nil
	}

	cpuMove := GenerateMove(*board)
	board.State[cpuMove] = board.CPUSymbol

	return nil
}
