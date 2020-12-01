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
	for _, v := range numbers {
		if v >= expectSum {
			continue
		}
		for _, j := range numbers {
			if v+j == expectSum {
				return v * j
			}
		}
	}

	return 0
}

func part2(numbers []int) int {
	for _, v := range numbers {
		if v >= expectSum {
			continue
		}
		for _, j := range numbers {
			if v+j >= expectSum {
				continue
			}
			for _, k := range numbers {
				if v+j+k == expectSum {
					return v * j * k
				}
			}
		}
	}

	return 0
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
	fmt.Printf("Part2: %v\n", part2(numbers))
}
