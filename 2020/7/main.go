package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EnriqueL8/adventofcode/utils"
)

type Bag struct {
	number int
	name   string
}

type Graph map[string][]Bag

const filename = "input.txt"

func traverse(bags []Bag, graph Graph, match, pMatch string) (counter int) {
	if len(bags) == 1 && bags[0].name == "other bags" {
		return
	}
	counter = 0
	for _, bag := range bags {
		if bag.name == match || bag.name == pMatch {
			return 1
		}
		if counter != 1 {
			bName := bag.name
			if bag.number == 1 {
				bName += "s"
			}

			if val, ok := graph[bName]; ok {
				if traverse(val, graph, match, pMatch) == 1 {
					return 1
				}
			}
		}
	}

	return counter
}

func calculateNumberOfBags(currentBag Bag, bags []Bag, graph Graph) int {
	n := 0
	for _, bag := range bags {
		n += bag.number
		bagName := bag.name
		if bag.number == 1 {
			bagName += "s"
		}
		if val, ok := graph[bagName]; ok {
			n += bag.number * calculateNumberOfBags(bag, val, graph)
		}
	}

	return n
}

func main() {
	lines, _ := utils.ReadLines("input.txt", "\n")
	graph := Graph{}
	for _, line := range lines {
		line = strings.Trim(line, ".")
		s1 := strings.Split(line, " contain ")
		if len(s1) < 2 {
			continue
		}
		key := s1[0]
		if _, ok := graph[key]; !ok {
			graph[key] = []Bag{}
		}

		bags := strings.Split(s1[1], ", ")
		for _, bag := range bags {
			s2 := strings.Split(bag, " ")
			number, _ := strconv.Atoi(string(s2[0]))
			name := strings.Join(s2[1:], " ")
			sBag := Bag{number: number, name: name}
			graph[key] = append(graph[key], sBag)
		}

	}

	totalCounter := 0
	match := "shiny gold bag"
	pMatch := "shiny gold bags"
	for key, value := range graph {
		if key == match || key == pMatch {
			continue
		}
		totalCounter += traverse(value, graph, match, pMatch)
	}

	fmt.Println(totalCounter)
	fmt.Println(calculateNumberOfBags(Bag{name: pMatch, number: 0}, graph[pMatch], graph))
}
