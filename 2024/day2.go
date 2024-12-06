package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day2part1() {
	// read input
	readFile, err := os.Open("day2input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var safeLevels = 0

	// iterate each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLineBySpaces := strings.Split(line, " ")

		// get first 2 numbers of the list
		firstNum, _ := strconv.ParseFloat(splitLineBySpaces[0], 64)
		secondNum, _ := strconv.ParseFloat(splitLineBySpaces[1], 64)

		// if difference is too large or numbers are the same, move on
		if math.Abs(firstNum-secondNum) > 3 || firstNum == secondNum {
			continue
		} else {
			// numbers are increasing
			if firstNum < secondNum {
				// add 1 to total if level was safe
				isSafe, _ := isLevelSafe(splitLineBySpaces, "increase")
				if isSafe {
					safeLevels++
				}
				// same shit as the one above
			} else if firstNum > secondNum {
				isSafe, _ := isLevelSafe(splitLineBySpaces, "decrease")
				if isSafe {
					safeLevels++
				}
			}
		}
	}
	fmt.Println(safeLevels)
}

func isLevelSafe(splitLineBySpaces []string, incOrDec string) (isSafe bool, brokenIdx int) {
	var safe bool
	var brokenIndex = 0
	fmt.Println("original:", splitLineBySpaces)

	firstNum, _ := strconv.ParseFloat(splitLineBySpaces[0], 64)
	secondNum, _ := strconv.ParseFloat(splitLineBySpaces[1], 64)
	if firstNum < secondNum {
		// iterate numbers starting at index 1 since 0 and 1 were just compared
		for i := 0; i < len(splitLineBySpaces)-1; i++ {
			// parse current & next number to floats
			currNum, _ := strconv.ParseFloat(splitLineBySpaces[i], 64)
			nextNum, _ := strconv.ParseFloat(splitLineBySpaces[i+1], 64)

			// if next is greater than current, difference is less than 3 and the numbers aren't the same
			// it is still safe
			fmt.Printf("next %f curr %f nextGreaterThanCurr %v\n", nextNum, currNum, nextNum > currNum)
			// fmt.Println(math.Abs(nextNum-currNum) <= 3)
			if nextNum > currNum && math.Abs(nextNum-currNum) <= 3 && nextNum != currNum {
				safe = true
				// if any checks fail, its not safe, break and move on
			} else {
				safe = false
			}
		}
	} else {
		for i := 0; i < len(splitLineBySpaces)-1; i++ {
			currNum, _ := strconv.ParseFloat(splitLineBySpaces[i], 64)
			nextNum, _ := strconv.ParseFloat(splitLineBySpaces[i+1], 64)
			fmt.Printf("next %f curr %f nextLessThanCurr %v\n", nextNum, currNum, nextNum < currNum)
			if nextNum < currNum && math.Abs(currNum-nextNum) <= 3 && nextNum != currNum {
				safe = true
			} else {
				safe = false
			}
		}
	}

	return safe, brokenIndex
}

func removeIndex(s []string, idx int) []string {
	newSlice := make([]string, len(s))
	copy(newSlice, s)
	if idx == 0 {
		return newSlice[1:]
	}
	return append(newSlice[:idx], newSlice[idx+1:]...)
}

func day2part2() {
	// read input
	readFile, err := os.Open("day2input.test")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var safeLevels = 0

	// iterate each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLineBySpaces := strings.Split(line, " ")
		// get first 2 numbers of the list
		firstNum, _ := strconv.ParseFloat(splitLineBySpaces[0], 64)
		secondNum, _ := strconv.ParseFloat(splitLineBySpaces[1], 64)

		// numbers are increasing
		if firstNum < secondNum {
			// add 1 to total if level was safe
			isSafe, _ := isLevelSafe(splitLineBySpaces, "increase")

			if isSafe {
				safeLevels++
			} else {
				// remove the bad number from the list
				// check if its safe now
				// withoutNext := removeIndex(splitLineBySpaces, brokenIdx)
				// withoutCurrent := removeIndex(splitLineBySpaces, brokenIdx-1)

				// isSafeWithoutNext, _ := isLevelSafe(withoutNext, "increase")
				// isSafeWithoutCurrent, _ := isLevelSafe(withoutCurrent, "increase")
				for i := 0; i < len(splitLineBySpaces); i++ {
					slicedLine := removeIndex(splitLineBySpaces, i)
					// fmt.Println("sliced:", slicedLine)
					isSafenow, _ := isLevelSafe(slicedLine, "increase")
					if isSafenow {
						safeLevels++
						break
					}
				}
				fmt.Println("-----------------")
			}
			// same shit as the one above
		} else if firstNum > secondNum {
			isSafe, _ := isLevelSafe(splitLineBySpaces, "decrease")
			if isSafe {
				safeLevels++
			} else {
				// remove current number and next number
				// remove curr to handle when first number is the bad one

				for i := 0; i < len(splitLineBySpaces); i++ {
					slicedLine := removeIndex(splitLineBySpaces, i)
					// fmt.Println("sliced:", slicedLine)
					isSafenow, _ := isLevelSafe(slicedLine, "decrease")
					if isSafenow {
						safeLevels++
						break
					}
				}
			}
		}
	}
	fmt.Println(safeLevels)
}
