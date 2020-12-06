package main

import (
	"fmt"
	"strings"

	"github.com/EnriqueL8/adventofcode/utils"
)

type Questions map[string]int

const filename = "input.txt"

// Need to use this because strins.Split returns an extra
// empty string "" as part of the slice
func getPeople(group string) []string {
	splitFn := func(c rune) bool {
		return c == '\n'
	}
	return strings.FieldsFunc(group, splitFn)
}

func updateQuestions(questions Questions, input string) {
	for _, question := range input {
		if question == '\n' {
			continue
		}
		sQuestion := string(question)
		if _, ok := questions[sQuestion]; ok {
			questions[sQuestion]++
		} else {
			questions[sQuestion] = 1
		}
	}
}

func part2(groups []string) int {
	totalAnsweredQuestions := 0
	for _, group := range groups {
		people := getPeople(group)
		if len(people) == 1 {
			totalAnsweredQuestions += len(people[0])
			continue
		}

		questions := Questions{}
		for _, person := range people {
			updateQuestions(questions, person)
		}

		// Delete the questions not answered by everyone
		for question, number := range questions {
			if number != len(people) {
				delete(questions, question)
			}
		}
		totalAnsweredQuestions += len(questions)
	}

	return totalAnsweredQuestions
}

func part1(groups []string) int {
	totalAnsweredQuestions := 0
	for _, group := range groups {
		questions := Questions{}
		updateQuestions(questions, group)
		totalAnsweredQuestions += len(questions)
	}

	return totalAnsweredQuestions
}

func main() {
	groups, _ := utils.ReadLines("input.txt", "\n\n")

	fmt.Printf("Part 1: %v\n", part1(groups))
	fmt.Printf("Part 2: %v\n", part2(groups))
}
