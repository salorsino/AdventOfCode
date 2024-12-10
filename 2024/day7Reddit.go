package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day7part1() {
	allGoals := []int{}
	allOperands := [][]int{}

	readFile, _ := os.Open("./day7input.txt")
	scan := bufio.NewScanner(readFile)
	for scan.Scan() {
		line := scan.Text()
		lineToks := strings.Split(line, ": ")

		goal, _ := strconv.Atoi(lineToks[0]) // Got goal!
		allGoals = append(allGoals, goal)

		operandsStrs := strings.Split(lineToks[1], " ")
		var operands []int
		for _, o := range operandsStrs { // Turn operands ([]string) into []int
			v, _ := strconv.Atoi(o)
			operands = append(operands, v)
		}
		allOperands = append(allOperands, operands)

		// log.Println("GOAL:", goal, "OPERANDS:", operands)
	}

	part1 := func() {
		total := 0
		for i, operands := range allOperands {
			subs := []int{operands[0]}
			fmt.Println(operands[1:])
			for _, o := range operands[1:] {
				var newSubs []int
				for _, s := range subs {
					newSubs = append(newSubs, s+o)

					if s == 0 {
						s = 1
					}
					newSubs = append(newSubs, s*o)
				}
				subs = newSubs
			}
			log.Println("SUBS:", subs)

			for _, s := range subs {
				if s == allGoals[i] {
					total += s
					break
				}
			}
		}

		log.Println("Part1:", total)
	}
	part1()

	part2 := func() {

		total := 0
		for i, operands := range allOperands {
			subs := []int{operands[0]}
			for _, o := range operands[1:] {
				var newSubs []int
				for _, s := range subs {
					newSubs = append(newSubs, s+o)

					if s == 0 {
						s = 1
					}
					newSubs = append(newSubs, s*o)

					// This is the only logic added to part 2... additional operation (||)
					ni, _ := strconv.Atoi(fmt.Sprintf("%d%d", s, o))
					newSubs = append(newSubs, ni)
				}
				subs = newSubs
			}
			// log.Println("SUBS:", subs)

			for _, s := range subs {
				if s == allGoals[i] {
					total += s
					break
				}
			}
		}

		log.Println("Part2:", total)
	}
	part2()
}
