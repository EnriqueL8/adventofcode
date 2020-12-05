package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

var columns = makeRange(0, 7)
var rows = makeRange(0, 127)

func binarySearch(values []int, input string, up, down string) int {
	mid := len(values) / 2
	if len(input) == 0 {
		return values[mid]
	}
	direction := string(input[0])
	switch {
	case len(values) == 0:
		return -1
	case direction == up:
		return binarySearch(values[:mid], input[1:], up, down)
	case direction == down:
		return binarySearch(values[mid:], input[1:], up, down)
	default:
		return values[mid]
	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	highestID := 0
	seats := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := binarySearch(rows, line[:7], "F", "B")
		column := binarySearch(columns, line[7:], "L", "R")
		id := row*8 + column
		if id > highestID {
			highestID = id
		}
		seats = append(seats, id)
	}
	sort.Ints(seats)

	fmt.Printf("Part 1: %v\n", highestID)
	for i, seat := range seats {
		if seat+1 != seats[i+1] {
			fmt.Printf("Part 2: %v\n", seat+1)
			break
		}
	}

}
