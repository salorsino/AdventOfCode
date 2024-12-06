package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day2part1() {
	readFile, err := os.Open("./day2Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	redMax := 12
	greenMax := 13
	blueMax := 14

	var gameIdSum = 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		onlyNumbers := regexp.MustCompile("[a-z A-Z:]")

		// get gameId string "Game 1:"
		gameString := strings.SplitAfterN(line, ":", 6)[0]

		// remove "Game" and ":"
		gameId, _ := strconv.Atoi(onlyNumbers.ReplaceAllString(gameString, ""))

		// remove "Game 1:" from line string
		line = strings.Replace(line, gameString, "", 1)

		// break games down
		games := strings.Split(line, ";")

		var isGamePossible = true
		// iterate over all the sets in the game
		for _, set := range games {

			// break the game into each set of colors
			colorBreakdown := strings.Split(set, ",")

			// iterate the set of colors
			for _, color := range colorBreakdown {
				// break out the color name and count
				colorCount := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorCount[1])
				color := colorCount[2]

				// if color in a set is greater than max, game is not possible
				if color == "red" && count > redMax {
					isGamePossible = false
				} else if color == "green" && count > greenMax {
					isGamePossible = false
				} else if color == "blue" && count > blueMax {
					isGamePossible = false
				}
			}
		}

		if isGamePossible {
			gameIdSum += gameId
		}
	}
	fmt.Println("Sum of game IDs:", gameIdSum)
}

func day2part2() {
	readFile, err := os.Open("./day2Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// total sum of all the powers of each game
	var powerSum = 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		// get gameId string "Game 1:"
		gameString := strings.SplitAfterN(line, ":", 6)[0]

		// remove "Game 1:" from line string
		line = strings.Replace(line, gameString, "", 1)

		// break games down
		games := strings.Split(line, ";")

		// declare vars for maximums
		var redMaximum, greenMaximum, blueMaximum = 0, 0, 0

		fmt.Println(gameString)

		// iterate over all the sets in the game
		for _, set := range games {

			// break the game into each set of colors
			colorBreakdown := strings.Split(set, ",")

			// iterate the set of colors
			for _, color := range colorBreakdown {
				// break out the color name and count
				colorCount := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorCount[1])
				color := colorCount[2]

				if color == "red" && count > redMaximum {
					redMaximum = count
				} else if color == "green" && count > greenMaximum {
					greenMaximum = count
				} else if color == "blue" && count > blueMaximum {
					blueMaximum = count
				}
			}
			fmt.Println(set)
		}
		powerSum += (redMaximum * greenMaximum * blueMaximum)
		fmt.Println("redMax:", redMaximum, "greenMax:", greenMaximum, "blueMax:", blueMaximum)
	}
	fmt.Println(powerSum)
}
