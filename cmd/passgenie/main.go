package main

import (
	"fmt"
	"os"
	"passgenie/internal/password"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	length := 12
	if envLength, exists := os.LookupEnv("PASSWORD_LENGTH"); exists {
		if l, err := strconv.Atoi(envLength); err == nil {
			length = l
		}
	}

	includeSymbols := false
	if envSymbols, exists := os.LookupEnv("INCLUDE_SYMBOLS"); exists {
		includeSymbols = envSymbols == "true"
	}

	gen := password.NewGenerator()
	pass, err := gen.Generate(length, includeSymbols)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated Password: %s\n", pass)
}
