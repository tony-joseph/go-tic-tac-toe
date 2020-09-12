package main

import (
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	number := GenerateRandomNumber(0, 10)
	if number >= 10 || number < 0 {
		t.Errorf("Number is not in range.")
	}

	number = GenerateRandomNumber(10, 20)
	if number >= 20 || number < 10 {
		t.Errorf("Number is not in range.")
	}
}

func TestIsFullLine(t *testing.T) {
	if IsFullLine(" ", " ", " ") {
		t.Errorf("Should be false when there is space in line.")
	}

	if IsFullLine("X", "O", "X") {
		t.Errorf("Should be false when first and second are not equal.")
	}

	if IsFullLine("X", "X", "O") {
		t.Errorf("Should be false when first and third are not equal.")
	}

	if !IsFullLine("X", "X", "X") {
		t.Errorf("Should be true when all are X.")
	}

	if !IsFullLine("O", "O", "O") {
		t.Errorf("Should be true when all are O.")
	}
}
