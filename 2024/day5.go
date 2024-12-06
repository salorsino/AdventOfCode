package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// check if slice contains a value
func contains(nums []int, val int) bool {
	for i := range nums {
		if nums[i] == val {
			return true
		}
	}
	return false
}

// return middle number of odd-length array
func getMiddle(input []string) string {
	return input[len(input)/2]
}

// func fixOrder(input []string, rules []string) []string {

// }

func day5part1() {

	// read inputs as two separate files
	readFile, _ := os.Open("day5inputpart1.txt")
	input, _ := os.Open("day5inputpart2.txt")

	defer readFile.Close()
	defer input.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	inputScanner := bufio.NewScanner(input)
	inputScanner.Split(bufio.ScanLines)

	/*
		declare map of rules where key is the number and val is slice of numbers that must be after it
		(not in numerical order)
		ex:
		75: [14, 15, 20, 90, 76]
	*/
	rulesMap := make(map[int][]int)

	// iterate all the rules from first input
	for fileScanner.Scan() {
		// split and get number & number for rule
		line := strings.Split(fileScanner.Text(), "|")
		ruleNum, _ := strconv.Atoi(line[0])
		restrictionNum, _ := strconv.Atoi(line[1])

		// append num to slice or declare slice in the map
		if len(rulesMap[ruleNum]) != 0 {
			rulesMap[ruleNum] = append(rulesMap[ruleNum], restrictionNum)
		} else {
			rulesMap[ruleNum] = []int{restrictionNum}
		}
	}

	// sort rules for each item in the map
	for k := range rulesMap {
		sort.Ints(rulesMap[k])
	}

	// var to store the valid rows
	var correctRows [][]string

	// iterate the rows to be checked
	for inputScanner.Scan() {
		// get current row to be checked
		inputLine := strings.Split(inputScanner.Text(), ",")

		// bool to check if row is in correct order
		var isCorrectOrder bool

		// iterate each num in the row
		for i, num := range inputLine {
			// copy the slice to split it up
			splitLine := make([]string, len(inputLine))
			copy(splitLine, inputLine)

			// get nums before & after current number
			beforeNum := splitLine[0:i]
			afterNum := splitLine[i+1:]

			// parse current to int for lookup
			numParsed, _ := strconv.Atoi(num)

			// sort lists for easier comparison
			sort.Strings(beforeNum)
			sort.Strings(afterNum)

			// if there are numbers before the one we are on
			// inspect to make sure they are not in the after ruleset
			if len(beforeNum) != 0 {
				for _, v := range beforeNum {
					// parse search num to int and search for it in rules
					parsedSearchNum, _ := strconv.Atoi(v)

					// if found in the rule map, row is not in correct order
					if contains(rulesMap[numParsed], parsedSearchNum) {
						isCorrectOrder = false
					} else {
						isCorrectOrder = true
					}

					// break from loop checking nums in beforeNum array
					if !isCorrectOrder {
						break
					}
				}

				// break from loop checking row since its already been declared broken
				if !isCorrectOrder {
					break
				}
			}
		}

		// push valid rows to results slice
		if isCorrectOrder {
			correctRows = append(correctRows, inputLine)
		}
	}

	// sum up middle numbers for each row
	middleSum := 0
	for i := range correctRows {
		middleNum, _ := strconv.Atoi(getMiddle(correctRows[i]))
		middleSum += middleNum
	}

	fmt.Println(middleSum)
}

func day5part2() {

	// read inputs as two separate files
	readFile, _ := os.Open("day5inputpart1.test")
	input, _ := os.Open("day5inputpart2.test")

	defer readFile.Close()
	defer input.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	inputScanner := bufio.NewScanner(input)
	inputScanner.Split(bufio.ScanLines)

	/*
		declare map of rules where key is the number and val is slice of numbers that must be after it
		(not in numerical order)
		ex:
		75: [14, 15, 20, 90, 76]
	*/
	rulesMap := make(map[int][]int)

	// iterate all the rules from first input
	for fileScanner.Scan() {
		// split and get number & number for rule
		line := strings.Split(fileScanner.Text(), "|")
		ruleNum, _ := strconv.Atoi(line[0])
		restrictionNum, _ := strconv.Atoi(line[1])

		// append num to slice or declare slice in the map
		if len(rulesMap[ruleNum]) != 0 {
			rulesMap[ruleNum] = append(rulesMap[ruleNum], restrictionNum)
		} else {
			rulesMap[ruleNum] = []int{restrictionNum}
		}
	}

	// sort rules for each item in the map
	for k := range rulesMap {
		sort.Ints(rulesMap[k])
	}

	// var to store the rows after they have been fixed
	var correctedRows [][]string
	var indexToMove int
	var indexOfNumInWrongOrder int

	// iterate the rows to be checked
	for inputScanner.Scan() {
		// get current row to be checked
		inputLine := strings.Split(inputScanner.Text(), ",")

		// bool to check if row is in correct order
		var isCorrectOrder bool

		// iterate each num in the row
		for i, num := range inputLine {
			// copy the slice to split it up
			splitLine := make([]string, len(inputLine))
			copy(splitLine, inputLine)

			// get nums before & after current number
			beforeNum := splitLine[0:i]
			afterNum := splitLine[i+1:]

			// parse current to int for lookup
			numParsed, _ := strconv.Atoi(num)

			// sort lists for easier comparison
			sort.Strings(beforeNum)
			sort.Strings(afterNum)

			// if there are numbers before the one we are on
			// inspect to make sure they are not in the after ruleset
			if len(beforeNum) != 0 {
				for _, v := range beforeNum {
					// parse search num to int and search for it in rules
					parsedSearchNum, _ := strconv.Atoi(v)

					// if found in the rule map, row is not in correct order
					if contains(rulesMap[numParsed], parsedSearchNum) {
						fmt.Println("rules for:", numParsed, rulesMap[numParsed])
						fmt.Printf("numParsed: %d - parsedSearchNum: %d\n", numParsed, parsedSearchNum)
						indexToMove = i
						indexOfNumInWrongOrder = slices.Index(inputLine, v)
						isCorrectOrder = false
						break
						// want to flip numParsed & parsedSearchNum in the original line
					} else {
						isCorrectOrder = true
					}

					// break from loop checking nums in beforeNum array
					if !isCorrectOrder {
						break
					}
				}

				// break from loop checking row since its already been declared as correct
				if !isCorrectOrder {
					break
				}
			}
		}

		// push valid rows to results slice
		if !isCorrectOrder {
			fmt.Println(inputLine)
			fmt.Println("indexToMove", indexToMove, "indexOfNumInWrongOrder:", indexOfNumInWrongOrder)
			correctedRows = append(correctedRows, inputLine)
		}
	}

	// sum up middle numbers for each row
	middleSum := 0
	for i := range correctedRows {
		middleNum, _ := strconv.Atoi(getMiddle(correctedRows[i]))
		middleSum += middleNum
	}

	fmt.Println(middleSum)
}
