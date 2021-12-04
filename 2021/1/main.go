package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const expectSum = 2020

func part1(numbers []int) int {
	increases := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1] > numbers[i] {
			increases++
		}
	}
	return increases
}

func part2(numbers []int, windowsize int) int {
	slidingWindowSums := []int{}
	for i := 0; i <= len(numbers)-windowsize+1; i++ {
		start := i
		end := start + windowsize
		numbersInWindow := numbers[start:end]
		windowSum := 0
		for _, number := range numbersInWindow {
			windowSum += number
		}
		slidingWindowSums = append(slidingWindowSums, windowSum)
	}

	return part1(slidingWindowSums)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	fmt.Printf("Part1: %v\n", part1(numbers))
	fmt.Printf("Part2: %v\n", part2(numbers, 3))
}
