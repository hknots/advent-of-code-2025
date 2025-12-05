package puzzles

import (
	"fmt"
	"os"
)

// Solver defines the function signature for a daily solution
type Solver func(file *os.File)

// Registry maps the day number to the solver function
var Registry = map[int]Solver{}

// Register adds a day's solution to the registry
func Register(day int, solver Solver) {
	if _, exists := Registry[day]; exists {
		panic(fmt.Sprintf("Day %d is already registered!", day))
	}
	Registry[day] = solver
}
