package main

import (
	"advent-of-code-2025/puzzles"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day number.")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day number")
		return
	}

	solver, exists := puzzles.Registry[day]
	if !exists {
		fmt.Printf("Day %d is not solved yet!\n", day)
		return
	}

	filename := fmt.Sprintf("input/day%02d.txt", day)
	cwd, _ := os.Getwd()
	fullPath := filepath.Join(cwd, filename)

	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Printf("--- Running Day %d ---\n", day)

	solver(file)
}
