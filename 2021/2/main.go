package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const expectSum = 2020

func calculateInstructionPart1(instruction string) (int, int) {
	s := strings.Split(instruction, " ")
	command := s[0]
	value, _ := strconv.Atoi(s[1])
	switch command {
	case "forward":
		return value, 0
	case "down":
		return 0, value
	case "up":
		return 0, -value
	}

	return 0, 0
}

func part1(instructions []string) int {
	horizontalPosition := 0
	depth := 0
	for _, instruction := range instructions {
		h, d := calculateInstructionPart1(instruction)
		horizontalPosition += h
		depth += d
	}

	return horizontalPosition * depth
}

func calculateInstructionPart2(instruction string, aim int) (int, int, int) {
	s := strings.Split(instruction, " ")
	command := s[0]
	value, _ := strconv.Atoi(s[1])
	switch command {
	case "forward":
		return value, aim * value, 0
	case "down":
		return 0, 0, value
	case "up":
		return 0, 0, -value
	}

	return 0, 0, 0
}

func part2(instructions []string) int {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, instruction := range instructions {
		h, d, a := calculateInstructionPart2(instruction, aim)
		horizontalPosition += h
		depth += d
		aim += a
	}

	return horizontalPosition * depth
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	fmt.Printf("Part1: %v\n", part1(instructions))
	fmt.Printf("Part2: %v\n", part2(instructions))
}
