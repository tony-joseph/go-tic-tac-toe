package main

import (
	"fmt"
	"testing"
)

func TestWinInNextMove(t *testing.T) {
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

	for _, line := range lines {
		board := Board{}
		board.Initialise("O")

		board.State[line[0]] = "X"
		if win, _ := WinInNextMove("X", board); win {
			t.Errorf("No win in next move")
		}

		board.State[line[1]] = "X"
		if win, move := WinInNextMove("X", board); !win || move != line[2] {
			t.Errorf("Win in next move")
		}

		board.State[line[0]] = "X"
		board.State[line[1]] = " "
		board.State[line[2]] = "X"
		if win, move := WinInNextMove("X", board); !win || move != line[1] {
			t.Errorf("Win in next move")
		}

		board.State[line[0]] = " "
		board.State[line[1]] = "X"
		board.State[line[2]] = "X"
		if win, move := WinInNextMove("X", board); !win || move != line[0] {
			t.Errorf("Win in next move")
		}
	}
}

func TestGenerateRandomMove(t *testing.T) {
	board := Board{}
	board.Initialise("O")

	board.State[0] = "X"
	board.State[1] = "X"
	board.State[2] = "X"

	move := GenerateRandomMove(board)

	if move < 3 || move > 8 {
		t.Errorf("Illegal move generation.")
	}
}

func TestGenerateMove(t *testing.T) {
	board := Board{}
	board.Initialise("O")

	// CPU win
	board.State[0] = "X"
	board.State[1] = "X"
	board.State[3] = "O"
	board.State[4] = "O"
	move := GenerateMove(board)
	if move != 5 {
		t.Errorf("Failed to generate CPU winning move.")
	}

	// Block player win
	board.State[0] = "X"
	board.State[1] = "X"
	board.State[3] = " "
	board.State[4] = " "
	move = GenerateMove(board)
	if move != 2 {
		t.Errorf("Failed to block player win.")
	}

	// Generate center move
	board.State[0] = " "
	board.State[1] = " "
	move = GenerateMove(board)
	if move != 4 {
		t.Errorf("Failed to generate center move.")
	}

	// Corner move
	board.State[4] = "X"
	move = GenerateMove(board)
	if move != 0 {
		t.Errorf("Invalid corner move.")
	}

	// Random move
	board.State[0] = "X"
	board.State[1] = "O"
	board.State[2] = "X"
	board.State[4] = "O"
	board.State[6] = "O"
	board.State[7] = "X"
	board.State[8] = "O"
	move = GenerateMove(board)
	fmt.Println(move)
	if !(move == 3 || move == 5) {
		t.Errorf("Invalid random move.")
	}
}

func TestGetEmptyCorner(t *testing.T) {
	board := Board{}
	board.Initialise("O")

	move, err := GetEmptyCorner(board)
	if err != nil {
		t.Errorf("Failed to generate corner move.")
	}
	if move != 0 {
		t.Errorf("Unexpected move generation.")
	}

	board.State[0] = "X"
	move, err = GetEmptyCorner(board)
	if err != nil {
		t.Errorf("Failed to generate corner move.")
	}
	if move != 2 {
		t.Errorf("Unexpected move generation.")
	}

	board.State[2] = "X"
	move, err = GetEmptyCorner(board)
	if err != nil {
		t.Errorf("Failed to generate corner move.")
	}
	if move != 6 {
		t.Errorf("Unexpected move generation.")
	}

	board.State[6] = "X"
	move, err = GetEmptyCorner(board)
	if err != nil {
		t.Errorf("Failed to generate corner move.")
	}
	if move != 8 {
		t.Errorf("Unexpected move generation.")
	}

	board.State[8] = "X"
	move, err = GetEmptyCorner(board)
	if err == nil {
		t.Errorf("Unexpected move generation.")
	}
}
