package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day3part1() {

	readFile, err := os.Open("./day3Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// declare 2d array to write schematic into
	var schematicMap [140][140]string

	// symbols to check for
	symbols := []string{
		"@",
		"/",
		"#",
		"$",
		"%",
		"&",
		"*",
		"-",
		"=",
		"+",
	}

	var lineIndex = 0

	// read input file and turn into 2d array
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for idx, char := range line {
			schematicMap[lineIndex][idx] = string(char)
		}
		lineIndex++
	}

	// iterate over 2d array
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			// fmt.Println(schematicMap[i][j])
			// check if we are a number

			var numString strings.Builder
			// var numSize int
			if num, err := strconv.ParseInt(schematicMap[i][j], 10, 32); err == nil {
				numString.WriteString(string(num))
				// check if the next two digits are numbers or not
				nextDigit, err := strconv.ParseInt(schematicMap[i][j+1], 10, 32)
				nextDigit2, err2 := strconv.ParseInt(schematicMap[i][j+2], 10, 32)
				fmt.Println(numString.String(), nextDigit, nextDigit2)
				if err != nil {
					numString.WriteString(string(nextDigit))
				}
				if err2 != nil {
					numString.WriteString(string(nextDigit2))
				}
				// fmt.Println("NUMBER:", numString)
			} else {
				fmt.Println(err)
			}
		}
	}
	fmt.Println(symbols)
}
