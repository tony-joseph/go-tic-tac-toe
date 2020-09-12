package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// ClearScreen clears output screen
func ClearScreen() {
	goos := runtime.GOOS
	if goos == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// GenerateRandomNumber generates a random number between limits
func GenerateRandomNumber(start int, limit int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limit-start) + start
}

// ReadMove reads and validate a move from stdin
func ReadMove() int {
	for {
		fmt.Println("Enter your move[1-9]:")
		var input string
		fmt.Scanln(&input)
		move, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid move. Please try again.")
			continue
		}
		return move - 1
	}
}

// IsFullLine checks if a winning line is complete
func IsFullLine(pos1 string, pos2 string, pos3 string) bool {
	if pos1 == " " {
		return false
	}

	if pos1 != pos2 {
		return false
	}

	if pos1 != pos3 {
		return false
	}

	return true
}
