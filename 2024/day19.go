package main

import (
	"fmt"
	"os"
	"strings"
)

var cache = make(map[string]bool)

var ans = false

func canMake(patterns []string, design string) (isValid bool) {
	// if we've already seen the pattern, return true
	// indicates it is valid
	if _, ok := cache[design]; ok {
		return cache[design]
	}

	// start assuming its false
	ans = false

	// if design we are checking is empty,
	// we've iterated from each char til the end and it was successful
	if len(design) == 0 {
		return true
	}

	// iterate through each pattern
	for _, pattern := range patterns {
		pattern = strings.Trim(pattern, " ")
		/*
			for each pattern, check if the string starts with it
			and
			call recursively check the string by removing the prefix we just checked

			if the loop never finds a design which as a pattern for a prefix
			ans will stay as false
			if it found a match but then next recursive call doesn't, canMake will return false
			for the previous iteration and loop will end

			ex:
			1. pattern = r, design = bggr
				move to next
			1. pattern = wr, design = bggr
				move to next
			1. pattern = b, design = bggr
				1. now make recursive call with ggr
					2. g as pattern, design = ggr
						2. recurse again with gr
							3. g = pattern, design = gr
								3. recurse again with r
									4. r is valid pattern
										4. recurse again with ""
											5. hits top if statement
											5. returns true
									4. canMake from "" returns true
										4. if statement passes
										4. ans = true
										4. mark "r" in cache as true
										4. return true
							3. canMake from "r" returns true
								3. if statement passes
								3. ans = true
								3. mark "gr" in cache as true
								3. return true
					2. canMake from "gr" returns true
						2. if statement passes
						2. ans = true
						2. mark "ggr" in cache as true
						2. return true
			1. canMake from "ggr" returns true
				1. if statement passes
				1. ans = true
				1. mark "bggr" in cache as true
				1. return true
		*/
		// fmt.Printf("Pattern: %s, Design: %s\n", string(pattern), design)
		if strings.HasPrefix(design, string(pattern)) {
			// fmt.Printf("Pattern: %s, Design: %s\n", string(pattern), design[len(string(pattern)):])
			isValid := canMake(patterns, design[len(string(pattern)):])
			if isValid {
				ans = true
			}
		}
	}
	cache[design] = true
	return ans
}

func day19part1() {
	data, err := os.ReadFile("./day19.test")
	if err != nil {
		fmt.Println("error opening input")
	}

	input := strings.Split(string(data), "\n\n")

	patterns, designs := input[0], input[1]
	patternSplit := strings.Split(patterns, ",")

	sum := 0
	for _, design := range strings.Split(designs, "\n") {
		if canMake(patternSplit, design) {
			sum++
		}
	}

	fmt.Println(sum)
}
