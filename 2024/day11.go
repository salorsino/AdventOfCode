package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(s []string, blinkNum int, dict map[string][]string) ([]string, int) {
	result := make([]string, 0)
	// cache := *dict
	sum := len(s)
	for _, v := range s {
		_, ok := dict[v]
		if ok {
			fmt.Println("cache hit")
			result = append(result, dict[v]...)
		} else if intVal, _ := strconv.Atoi(v); intVal == 0 {
			result = append(result, "1")
		} else if len(v)%2 == 0 {
			firstHalf, _ := strconv.Atoi(v[:int(float64(len(v)/2))])
			secondHalf, _ := strconv.Atoi(v[int(float64(len(v)/2)):])

			result = append(result, strconv.Itoa(firstHalf), strconv.Itoa(secondHalf))
			sum++
		} else {
			valAsInt, _ := strconv.Atoi(v)
			result = append(result, strconv.Itoa(valAsInt*2024))
		}
	}
	return result, sum
}

func blink2(s string, blinkNum int, dict map[string][]int) int {
	if s == "0" {
		return 1
	} else if len(s)%2 == 0 {
		// result := 0
		// firstHalf, _ := strconv.Atoi(s[:int(float64(len(s)/2))])
		// secondHalf, _ := strconv.Atoi(s[int(float64(len(s)/2)):])

		// result += blink2(firstHalf, blinkNum-1, dict)
	}
	return 1
}

func day11part2() {
	data, err := os.ReadFile("./day11input.test")
	if err != nil {
		fmt.Println("error opening input")
	}

	input := strings.Split(string(data), " ")
	result := make([]string, len(input))
	// cache := make(map[string][]int)
	sum := 0
	// for _, v := range input {
	// 	// sum, _ = blink2(v, 75, cache)
	// }

	fmt.Println(len(result)-len(input), sum)

}
func day11part1() {
	data, err := os.ReadFile("./day11input.test")
	if err != nil {
		fmt.Println("error opening input")
	}

	input := strings.Split(string(data), " ")
	result := make([]string, len(input))
	cache := make(map[string][]string)
	sum := 0
	for _, v := range input {
		localResult := []string{v}
		localSum := 0
		for i := 0; i < 6; i++ {
			// fmt.Println(localResult)
			fmt.Println("blink:", i)
			curr, currSum := blink(localResult, i, cache)
			localResult = curr
			localSum = currSum
			// cache[v] = curr
		}
		sum += localSum
	}

	fmt.Println(len(result)-len(input), sum)

}
