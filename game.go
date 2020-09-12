package main

import (
	"errors"
)

// WinInNextMove checks if player can win in next move
func WinInNextMove(symbol string, board Board) (bool, int) {
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
		if board.State[line[0]] == symbol && board.State[line[1]] == symbol && board.State[line[2]] == " " {
			return true, line[2]
		}

		if board.State[line[0]] == symbol && board.State[line[1]] == " " && board.State[line[2]] == symbol {
			return true, line[1]
		}

		if board.State[line[0]] == " " && board.State[line[1]] == symbol && board.State[line[2]] == symbol {
			return true, line[0]
		}
	}
	return false, 0
}

// GenerateRandomMove finds a random move from possible moves
func GenerateRandomMove(board Board) int {
	legalMoves := make([]int, 0)
	for i, player := range board.State {
		if player == " " {
			legalMoves = append(legalMoves, i)
		}
	}
	moveIndex := GenerateRandomNumber(0, len(legalMoves))
	return legalMoves[moveIndex]
}

// GetEmptyCorner checks if corner move is possible
func GetEmptyCorner(board Board) (int, error) {
	corners := [4]int{0, 2, 6, 8}

	for _, corner := range corners {
		if board.State[corner] == " " {
			return corner, nil
		}
	}

	return 0, errors.New("No corner moves found")
}

// GenerateMove finds the most optimal move
func GenerateMove(board Board) int {
	// Make winning move
	if win, move := WinInNextMove(board.CPUSymbol, board); win {
		return move
	}

	// Block opponent's winning move
	if win, move := WinInNextMove(board.PlayerSymbol, board); win {
		return move
	}

	// Make center move if possible.
	if board.State[4] == " " {
		return 4
	}

	// Make corner move if possible
	corner, err := GetEmptyCorner(board)
	if err == nil {
		return corner
	}

	return GenerateRandomMove(board)
}
