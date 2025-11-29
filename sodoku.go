package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// 1. Check for exactly 10 arguments (1 program name + 9 rows)
	if len(os.Args) != 10 { Error(); return }
	g := [9][9]int{}

	// 2. Parse and Validate Input
	for r, arg := range os.Args[1:] {
		if len(arg) != 9 { Error(); return } // Row length check
		for c, ch := range arg {
			if ch >= '1' && ch <= '9' {
				n := int(ch - '0')
				// Check if the starting number conflicts with existing numbers
				if !isValid(g, r, c, n) { Error(); return }
				g[r][c] = n
			} else if ch != '.' { 
				// Invalid character check
				Error(); return 
			}
		}
	}

	// 3. Solve and Print
	if solve(&g) {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				z01.PrintRune(rune(g[r][c] + '0'))
				if c != 8 { z01.PrintRune(' ') }
			}
			z01.PrintRune('\n')
		}
	} else { Error() }
}

// Helper: Checks Row, Column, and Box in one loop
func isValid(g [9][9]int, r, c, n int) bool {
	for i := 0; i < 9; i++ {
		if g[r][i] == n || g[i][c] == n || g[r/3*3+i/3][c/3*3+i%3] == n {
			return false
		}
	}
	return true
}

// Helper: Backtracking Solver
func solve(g *[9][9]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == 0 { // Find empty cell
				for n := 1; n <= 9; n++ {
					if isValid(*g, r, c, n) {
						g[r][c] = n // Place number
						if solve(g) { return true } // Recurse
						g[r][c] = 0 // Backtrack
					}
				}
				return false // No number worked here
			}
		}
	}
	return true // No empty cells left
}

// Helper: Print Error
func Error() { for _, r := range "Error\n" { z01.PrintRune(r) } }
