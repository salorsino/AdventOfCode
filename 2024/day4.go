package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func checkXmasHorizontal(s string) int {
	if s == "XMAS" {
		return 1
	}
	return 0
}

/*
As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt;
she'd like to know if you could help her with her word search (your puzzle input).
She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards,
or even overlapping other words. It's a little unusual, though, as you don't merely need to
find one instance of XMAS - you need to find all of them.
*/
func day4part1() {
	readFile, err := os.Open("day4input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// create 2d grid
	grid := make([][]string, 140)
	lineNum := 0

	// populate grid with characters from input
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid[lineNum] = make([]string, 140)
		for idx, char := range line {
			grid[lineNum][idx] = string(char)
		}
		lineNum++
	}

	xmasSum := 0

	// iterate grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// get current row & letter
			currentRow := grid[i]
			curr := currentRow[j]

			// reverse the row to check for backwards
			reversedRow := make([]string, len(currentRow))
			copy(reversedRow, currentRow)
			slices.Reverse(reversedRow)
			currReversed := reversedRow[j]

			// check horizontals in forward direction
			if j+3 < len(currentRow) && curr == "X" {
				xmasSum += checkXmasHorizontal(string(strings.Join(currentRow[j:j+4], "")))
			}
			// check horizontal in backwards direction
			if j+3 < len(reversedRow) && currReversed == "X" {
				xmasSum += checkXmasHorizontal(string(strings.Join(reversedRow[j:j+4], "")))
			}
			// check vertical going down
			if i+3 < len(grid) && curr == "X" {
				s := curr + grid[i+1][j] + grid[i+2][j] + grid[i+3][j]
				xmasSum += checkXmasHorizontal(s)
			}
			// check veritcal going up
			if i-3 >= 0 && curr == "X" {
				s := curr + grid[i-1][j] + grid[i-2][j] + grid[i-3][j]
				xmasSum += checkXmasHorizontal(s)
			}
			// check downward to the right diagonals
			if j+3 < len(currentRow) && i+3 < len(grid) && curr == "X" {
				s := curr + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3]
				xmasSum += checkXmasHorizontal(s)
			}
			// check downward to the left diagonals
			if j-3 >= 0 && i+3 < len(grid) && curr == "X" {
				s := curr + grid[i+1][j-1] + grid[i+2][j-2] + grid[i+3][j-3]
				xmasSum += checkXmasHorizontal(s)
			}
			// check upward to the right diagonals
			if j+3 < len(currentRow) && i-3 >= 0 && curr == "X" {
				s := curr + grid[i-1][j+1] + grid[i-2][j+2] + grid[i-3][j+3]
				xmasSum += checkXmasHorizontal(s)
			}
			// check downward to the left diagonals
			if j-3 >= 0 && i-3 >= 0 && curr == "X" {
				s := curr + grid[i-1][j-1] + grid[i-2][j-2] + grid[i-3][j-3]
				xmasSum += checkXmasHorizontal(s)
			}
		}
	}
	fmt.Println(xmasSum)
}

func day4part2() {
	readFile, err := os.Open("day4input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// create 2d grid
	grid := make([][]string, 140)
	lineNum := 0

	// populate grid with characters from input
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid[lineNum] = make([]string, 140)
		for idx, char := range line {
			grid[lineNum][idx] = string(char)
		}
		lineNum++
	}

	xmasSum := 0

	// iterate grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// get current row & letter
			currentRow := grid[i]
			curr := currentRow[j]

			// check that we are still in bounds
			if i-1 >= 0 && i+1 < len(grid) && j-1 >= 0 && j+1 < len(currentRow) && curr == "A" {
				// get letters at each index
				topLeft := grid[i-1][j-1]
				topRight := grid[i-1][j+1]
				bottomRight := grid[i+1][j+1]
				bottomLeft := grid[i+1][j-1]

				// check possible combinations
				if topLeft == "M" && topRight == "S" && bottomLeft == "M" && bottomRight == "S" {
					xmasSum++
				} else if topLeft == "S" && topRight == "M" && bottomLeft == "S" && bottomRight == "M" {
					xmasSum++
				} else if topLeft == "S" && topRight == "S" && bottomLeft == "M" && bottomRight == "M" {
					xmasSum++
				} else if topLeft == "M" && topRight == "M" && bottomLeft == "S" && bottomRight == "S" {
					xmasSum++
				}
			}
		}
	}
	fmt.Println(xmasSum)
}
