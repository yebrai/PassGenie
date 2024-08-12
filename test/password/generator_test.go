package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	gen := NewGenerator()

	length := 12
	includeSymbols := true

	pass, err := gen.Generate(length, includeSymbols)
	if err != nil {
		t.Fatalf("Error generating password: %v", err)
	}
	if len(pass) != length {
		t.Fatalf("Expected password length %d, got %d", length, len(pass))
	}

	// Check if password contains symbols
	if includeSymbols && !containsSymbols(pass) {
		t.Fatalf("Password does not contain symbols")
	}
}

func containsSymbols(password string) bool {
	symbolChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	for _, char := range password {
		if contains(symbolChars, char) {
			return true
		}
	}
	return false
}

func contains(chars string, char rune) bool {
	for _, c := range chars {
		if c == char {
			return true
		}
	}
	return false
}
