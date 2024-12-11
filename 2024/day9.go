package main

import (
	"fmt"
	"os"
	"sort"
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
	// var inputArray []string

	// make a map of the starting index of each free space, and the length
	spacesMap := make(map[int]int)

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
		if i%2 != 0 {
			spacesMap[len(strings.ReplaceAll(newString.String(), ",", ""))] = lengthToFill
		}
		if numToFill > 9 && i%2 == 0 {
			newS := strconv.Itoa(numToFill)
			newString.WriteString(strings.Repeat(newS+",", lengthToFill))
		} else {
			j := 0
			for j < lengthToFill {
				if i%2 == 0 {
					newS := strconv.Itoa(numToFill)
					newString.WriteString(newS + ",")
				} else {
					newString.WriteString(".,")
				}
				j++
			}
		}
	}

	// get string value
	sVal := strings.Split(newString.String(), ",")

	var spaceIndicesSorted []int
	// get list of free spaces
	for idx := range spacesMap {
		spaceIndicesSorted = append(spaceIndicesSorted, idx)
	}
	// sort the ints
	sort.Ints(spaceIndicesSorted)
	j := len(sVal) - 2
	// iterate final string at indicies of free spaces
	for _, idx := range spaceIndicesSorted {
		if !strings.Contains(strings.Join(sVal, ""), ".") {
			strings.Join(sVal, "")
			break
		}
		lengthOfSpaces := spacesMap[idx]
		start := idx
		// iterate the free spaces we are currently on
		for i := 0; i < lengthOfSpaces; i++ {
			// make sure we are not at "."
			for sVal[j] == "." {
				sVal = sVal[:j]
				j--
			}
			if start >= len(sVal) {
				break
			}
			fmt.Printf("leftVal: %s rightVal: %s\n", sVal[start], sVal[j])
			sVal[start], sVal[j] = sVal[j], sVal[start]

			sVal = sVal[:j]
			j--
			start++
		}
	}

	// iterate once more to sum everything up
	sum := 0
	for i, s := range sVal {
		result, _ := strconv.Atoi(s)
		sum += int(i) * result
	}

	fmt.Println(sVal)

	fmt.Println(sum)

}
