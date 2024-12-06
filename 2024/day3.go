package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isMultiplication(section string) bool {
	return !strings.HasPrefix(section, "do")
}

func multiplyNumbers(section string) int {

	// get just the numbers
	re := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	numsOnly := strings.Split(re.FindString(section), ",")

	// convert & multiply
	first, _ := strconv.Atoi(numsOnly[0])
	second, _ := strconv.Atoi(numsOnly[1])
	return first * second
}

func day3part1() {
	data, err := os.ReadFile("./day3input.txt")
	if err != nil {
		fmt.Println("error opening input")
	}

	fileContent := string(data)

	// create a slice of just the valid mul() commands
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	validSections := re.FindAllString(fileContent, -1)

	var sum int

	for x := range validSections {
		section := validSections[x]
		sum += multiplyNumbers(section)
	}
	fmt.Println(sum)
}

func day3part2() {
	data, err := os.ReadFile("./day3input.txt")
	if err != nil {
		fmt.Println("error opening input")
	}

	fileContent := string(data)

	// create a slice of just the valid mul() commands
	re := regexp.MustCompile(`(do\(\)|don\'t\(\)|mul\(\d{1,3},\d{1,3}\))`)
	validSections := re.FindAllString(fileContent, -1)

	var sum int

	// start enabled
	doOrDont := "do"

	for i := 0; i < len(validSections); i++ {
		section := validSections[i]

		// set doOrDont based on most recent we see
		if !isMultiplication(section) {
			doOrDont = strings.Replace(section, "()", "", 1)
			// if we are on mul() section and do is enabled, sum
		} else if doOrDont == "do" {
			sum += multiplyNumbers(section)
		}
	}
	fmt.Println(sum)
}
