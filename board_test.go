package main

import (
	"testing"
)

func TestInitialise(t *testing.T) {
	// Test symbol X
	board := Board{}
	board.Initialise("X")

	if board.CPUSymbol != "X" {
		t.Errorf("CPU is not X")
	}

	if board.PlayerSymbol != "O" {
		t.Errorf("Player is not O")
	}

	symbolCount := 0
	for _, item := range board.State {
		if item == "X" {
			symbolCount++
		} else if item != " " {
			t.Errorf("Initial state is not space( )")
		}
	}
	if symbolCount != 1 {
		t.Errorf("One cell should be X")
	}

	// Test symbol O
	board = Board{}
	board.Initialise("O")

	if board.CPUSymbol != "O" {
		t.Errorf("CPU is not O")
	}

	if board.PlayerSymbol != "X" {
		t.Errorf("Player is not X")
	}

	for _, item := range board.State {
		if item != " " {
			t.Errorf("Initial state is not space( )")
		}
	}
}

func TestIsValidMove(t *testing.T) {
	board := Board{}
	board.Initialise("O")

	if board.IsValidMove(9) {
		t.Errorf("Moves greater than 8 should not be allowed.")
	}

	if board.IsValidMove(-1) {
		t.Errorf("Moves less than 0 should not be allowed.")
	}

	// All moves are allowed in empty board
	for i := 0; i < 9; i++ {
		if !board.IsValidMove(i) {
			t.Errorf("All moves should be allowed in empty board.")
		}
	}

	// Moves should not be allowed in filled cell
	board.State[4] = "X"
	if board.IsValidMove(4) {
		t.Errorf("Move should not be allowed in filled cell.")
	}
}

func TestMove(t *testing.T) {
	board := Board{}
	board.Initialise("O")

	err := board.Move(1)
	if err != nil {
		t.Errorf("Move should be allowed.")
	}

	xCount := 0
	for _, symbol := range board.State {
		if symbol == "X" {
			xCount++
		}
	}
	if xCount != 1 {
		t.Errorf("X count should be 1.")
	}

	err = board.Move(1)
	if err == nil {
		t.Errorf("Move to filled cell should not be allowed.")
	}

	// No new moves when there is a winner
	board = Board{}
	board.Initialise("0")
	board.State[0] = "X"
	board.State[1] = "X"
	board.Move(2)

	for i := 3; i < 9; i++ {
		if board.State[i] != " " {
			t.Errorf("No moves should be made when there is already a winner.")
		}
	}
}

func TestGetWinner(t *testing.T) {
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

	for i := 0; i < len(lines); i++ {
		board := Board{}
		board.Initialise("O")

		board.State[lines[i][0]] = "X"
		board.State[lines[i][1]] = "X"
		if hasWinner, _ := board.GetWinner(); hasWinner {
			t.Errorf("Should be no-winner")
		}

		board.State[lines[i][2]] = "O"
		if hasWinner, _ := board.GetWinner(); hasWinner {
			t.Errorf("Should be no-winner")
		}

		board.State[lines[i][2]] = "X"
		if hasWinner, winner := board.GetWinner(); !hasWinner || winner != "X" {
			t.Errorf("Should be X win")
		}
	}

	// Test Draw
	board := Board{}
	board.Initialise("O")

	board.State[0] = "X"
	board.State[1] = "X"
	board.State[2] = "O"
	board.State[3] = "O"
	board.State[4] = "O"
	board.State[5] = "X"
	board.State[6] = "X"
	board.State[7] = "X"
	board.State[8] = "O"

	hasWinner, winner := board.GetWinner()
	if !hasWinner || winner != " " {
		t.Errorf("Should be draw.")
	}
}
