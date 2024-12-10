package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day9part1() {

	data, err := os.ReadFile("./day9input.test")
	if err != nil {
		fmt.Println("error opening input")
	}

	inputString := string(data)

	var newString strings.Builder

	// build string based on file input
	/*
		at 0 and even indexes and the size of the blocks
		odd indexes are sizes of free space
		build string by concatenating strings with length inputString[i] and value floor(i/2) starting at i
		if i%2 != 0 then add "." instead of i%2
		ex:
		12345 == 0..111....22222
		starting at index 0 (i), fill 1 value (inputString[i] = 1) with 0s (floor(i/2))
		i == 1, 1 % 2 != 0, so append/fill 2 values with "."
		i == 2 	2%2 == 0, 2 / 2 == 1, inputString[i] = 3, so fill 3 values with 1
	*/
	for i, s := range inputString {
		numToFill := int(i / 2)
		lengthToFill, _ := strconv.Atoi(string(s))

		j := 0
		for j < lengthToFill {
			if i%2 == 0 {
				newS := strconv.Itoa(numToFill)
				newString.WriteString(newS)
			} else {
				newString.WriteString(".")
			}
			j++
		}
	}

	// get string value
	sVal := newString.String()

	// iterate string and move numbers to free spaces
	for i, s := range sVal {
		j := len(sVal) - 1
		for string(s) == "." {

		}
	}

	fmt.Println(newString.String())
}
