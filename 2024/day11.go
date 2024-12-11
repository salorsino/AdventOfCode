package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(s []string) []string {
	result := make([]string, 0)
	for _, v := range s {
		if intVal, _ := strconv.Atoi(v); intVal == 0 {
			result = append(result, "1")
		} else if len(v)%2 == 0 {
			firstHalf, _ := strconv.Atoi(v[:int(float64(len(v)/2))])
			secondHalf, _ := strconv.Atoi(v[int(float64(len(v)/2)):])

			result = append(result, strconv.Itoa(firstHalf), strconv.Itoa(secondHalf))
		} else {
			valAsInt, _ := strconv.Atoi(v)
			result = append(result, strconv.Itoa(valAsInt*2024))
		}
	}
	return result
}

func day11part1() {
	data, err := os.ReadFile("./day11input.test")
	if err != nil {
		fmt.Println("error opening input")
	}

	input := strings.Split(string(data), " ")
	result := make([]string, len(input))
	copy(result, input)
	for i := 0; i < 25; i++ {
		result = blink(result)
	}

	fmt.Println(len(result))

}
