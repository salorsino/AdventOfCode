package main

import (
	"bufio"
	"fmt"
	"os"
)

func day6part1() {
	readFile, err := os.Open("day6input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// create 2d grid
	grid := make([][]string, 130)
	lineNum := 0

	// map to track where we have been already
	// key is XY combined into string
	visited := make(map[string]bool)

	// track and direction index of arrow
	arrowIdx := []int{0, 0}
	arrowDirection := ""

	// populate grid with characters from input
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid[lineNum] = make([]string, 130)
		for idx, char := range line {
			pointer := string(char)
			if pointer == "^" || pointer == ">" || pointer == "<" || pointer == "v" {
				arrowIdx = []int{lineNum, idx}
				arrowDirection = pointer
			}
			grid[lineNum][idx] = string(char)
		}
		lineNum++
	}

	exitFound := false

	upperBound := 0
	leftBound := 0
	rightBound := len(grid[0])
	bottomBound := len(grid)

	i := arrowIdx[0]
	j := arrowIdx[1]

	// check which direction the arrow is facing and iterate that way
	for !exitFound {
		if arrowDirection == "^" {
			// while still in bounds going up
			for i >= upperBound {
				if i-1 < upperBound {
					exitFound = true
					break
				}
				curr := grid[i-1][j]
				// if next char is ".", save coordinates to map && move up
				// only save coordinates if we haven't been there before
				if curr == "#" {
					arrowDirection = ">"
					break
				} else {
					coordinates := string(i-1) + string(j)
					if !visited[coordinates] {
						visited[coordinates] = true
					}
					i--
				}
			}

			if exitFound {
				break
			}
		} else if arrowDirection == ">" {
			// while still in bounds going right
			for j <= rightBound {
				if j+1 >= rightBound {
					exitFound = true
					break
				}
				// if next char is ".", save coordinates to map && move up
				// only save coordinates if we haven't been there before

				curr := grid[i][j+1]
				if curr == "#" {
					arrowDirection = "v"
					break
				} else {
					coordinates := string(i) + string(j+1)
					if !visited[coordinates] {
						visited[coordinates] = true
					}
					j++
				}
			}
			if exitFound {
				break
			}
		} else if arrowDirection == "v" {
			// while still in bounds going right
			for i <= bottomBound {
				// if next move will be out of bounds, we made it out of the grid
				if i+1 >= bottomBound {
					exitFound = true
					break
				}
				// if next char is ".", save coordinates to map && move up
				// only save coordinates if we haven't been there before

				curr := grid[i+1][j]
				if curr == "#" {
					arrowDirection = "<"
					break
				} else {
					coordinates := string(i+1) + string(j)
					if !visited[coordinates] {
						visited[coordinates] = true
					}
					i++
				}
			}
			if exitFound {
				break
			}
		} else if arrowDirection == "<" {
			// while still in bounds going right
			for j >= leftBound {
				if j-1 < leftBound {
					exitFound = true
					break
				}
				// if next char is ".", save coordinates to map && move up
				// only save coordinates if we haven't been there before

				curr := grid[i][j-1]
				if curr == "#" {
					arrowDirection = "^"
					break
				} else {
					coordinates := string(i) + string(j-1)
					if !visited[coordinates] {
						visited[coordinates] = true
					}
					j--
				}
			}
			if exitFound {
				break
			}
		}
	}

	fmt.Println(len(visited))
}
