package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
Reads input file
removes non digit characters
creates 2 digit number from the first & last digit
adds each 2 digit number to a sum
*/
func day1part1() {
	readFile, err := os.Open("Day1Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		calibration := fileScanner.Text()

		// remove all non-number characters
		regEx := regexp.MustCompile("[a-zA-Z]")

		justNumbers := regEx.ReplaceAllString(calibration, "")

		first := string(justNumbers[0])
		second := string(justNumbers[len(justNumbers)-1])

		num, _ := strconv.Atoi(first + second)
		sum += num
	}
	fmt.Println("Part 1 Total:", sum)
}

/*
Same as part 1 except first has to convert word numbers to ints
After conversion, take first number and last number and add to sum
*/
func day1part2() {
	readFile, err := os.Open("./Day1Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	type numberAsWord struct {
		word        string
		numberValue string
	}

	numsAsWords := []numberAsWord{
		{word: "twone", numberValue: "1"},
		{word: "one", numberValue: "1"},
		{word: "two", numberValue: "2"},
		{word: "three", numberValue: "3"},
		{word: "four", numberValue: "4"},
		{word: "five", numberValue: "5"},
		{word: "six", numberValue: "6"},
		{word: "seven", numberValue: "7"},
		{word: "eight", numberValue: "8"},
		{word: "nine", numberValue: "9"},
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		calibrationString := fileScanner.Text()

		var replaced = calibrationString

		// replace words with numeric value
		for _, numWord := range numsAsWords {
			replaced = strings.ReplaceAll(replaced, numWord.word, numWord.numberValue)
		}

		// remove all non-number characters
		regEx := regexp.MustCompile("[a-zA-Z]")
		justNumbers := regEx.ReplaceAllString(replaced, "")

		// get first & last digits
		first := string(justNumbers[0])
		second := string(justNumbers[len(justNumbers)-1])

		// build 2 digit number from first & last digits
		num, _ := strconv.Atoi(first + second)
		fmt.Println("original:", calibrationString, "just numbers:", justNumbers, "twoDigit:", num)
		sum += num
	}

	fmt.Println("Part 2 Total:", sum)
}
