package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Register(5, SolveDay05)
}

type freshRange struct {
	Min int
	Max int
}

// TODO: Determine if we should handle error handling here or in main.go

// SolveDay05 Answers how many fresh fruits there are and how many unique fresh ids there are
// Ref: https://adventofcode.com/2025/day/5
func SolveDay05(file *os.File) {
	scanner := bufio.NewScanner(file)
	freshFruits := make([]int, 0, 100)
	freshRanges := make([]freshRange, 0, 100)
	scanningRanges := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" { // Empty whitespace to split freshRanges and values
			scanningRanges = false
			freshRanges = mergeFreshRanges(freshRanges)
			continue
		}

		if scanningRanges {
			parts := strings.SplitN(line, "-", 2)
			partsLen := len(parts)

			if partsLen != 2 {
				fmt.Printf("invalid length of parts: %d is expected to be 2\n", partsLen)
				return
			}

			fr, err := createFreshRange(parts)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			freshRanges = append(freshRanges, fr)
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("failed to convert value: \"%s\" to an integer\n", line)
				return
			}

			for _, fr := range freshRanges {
				if value >= fr.Min && value <= fr.Max {
					freshFruits = append(freshFruits, value)
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Printf("The number of fresh fruits: %d\n", len(freshFruits))
	fmt.Printf("The frequency of fresh fruit ids are: %d\n", calculateIdFrequency(freshRanges))
}

// mergeFreshRanges merges freshRange's based on their overlapping values
func mergeFreshRanges(frs []freshRange) []freshRange {
	mergedFreshRanges := make([]freshRange, 0, 50)

	sort.Slice(frs, func(i, j int) bool {
		return frs[i].Min < frs[j].Min
	})

	current := frs[0]

	for _, next := range frs[1:] {
		if next.Min <= current.Max+1 { // + 1 to include adjacent (e.g., 1-3 and 4-6 should merge to 1-6)
			current.Max = max(current.Max, next.Max)
		} else {
			mergedFreshRanges = append(mergedFreshRanges, current)
			current = next
		}
	}

	// Append the final excluded range
	return append(mergedFreshRanges, current)
}

func createFreshRange(parts []string) (freshRange, error) {
	minValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return freshRange{}, fmt.Errorf("failed to parse Min value (index 0): %w", err)
	}

	maxValue, err := strconv.Atoi(parts[1])
	if err != nil {
		return freshRange{}, fmt.Errorf("failed to parse Max value (index 1): %w", err)
	}

	return freshRange{
		Min: minValue,
		Max: maxValue,
	}, nil
}

func calculateIdFrequency(frs []freshRange) int {
	var count int

	for _, fr := range frs {
		count += (fr.Max - fr.Min) + 1 // + 1 to include starting id
	}

	return count
}
