package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1part1() {
	// read input
	readFile, err := os.Open("day1input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// arrays to hold each list
	var firstList []int
	var secondList []int

	// iterate each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Split(line, "   ")

		firstNum, _ := strconv.Atoi(splitLine[0])
		secondNum, _ := strconv.Atoi(splitLine[1])

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)
	}

	// sort both lists
	sort.Ints(firstList)
	sort.Ints(secondList)

	var differenceSum = 0
	for x := range firstList {

		firstFloat := float64(firstList[x])
		secondFloat := float64(secondList[x])
		diff := int(math.Abs(firstFloat - secondFloat))
		fmt.Printf("First %d second %d difference %d\n", firstList[x], secondList[x], diff)
		differenceSum += diff
	}

	fmt.Println(differenceSum)
}

func day1part2() {
	// read input
	readFile, err := os.Open("day1input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// map to show if number exists
	var firstList []int
	secondList := make(map[int]int)

	// iterate each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Split(line, "   ")

		firstNum, _ := strconv.Atoi(splitLine[0])
		secondNum, _ := strconv.Atoi(splitLine[1])

		firstList = append(firstList, firstNum)
		secondList[secondNum] += 1
	}

	var similarityScore = 0

	for _, num := range firstList {
		if secondList[num] != 0 {
			// fmt.Printf("First %d is in second: %v\n", num, secondList[num])
			similarityScore += num * secondList[num]
		}
	}

	fmt.Println(similarityScore)
}
