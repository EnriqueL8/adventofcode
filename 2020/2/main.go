package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(line string) bool {
	lowerRange, upperRange, letter, word := split(line)
	freq := map[string]int{}
	for _, char := range word {
		sChar := string(char)
		if val, ok := freq[sChar]; ok {
			freq[sChar] = val + 1
		} else {
			freq[sChar] = 1
		}
	}

	if lowerRange <= freq[letter] && freq[letter] <= upperRange {
		return true
	}

	return false
}

func part2(line string) bool {
	firstPos, secondPos, letter, word := split(line)
	letterAtFirstPos := string(word[firstPos-1])
	letterAtSecondPos := string(word[secondPos-1])
	// XOR
	return (letterAtFirstPos == letter && letterAtSecondPos != letter) || (letterAtFirstPos != letter && letterAtSecondPos == letter)
}

func split(line string) (firstNumber int, secondNumber int, letter string, word string) {
	s := strings.Split(line, ":")
	if len(s) != 2 {
		return
	}
	parms := strings.Split(s[0], " ")
	if len(parms) != 2 {
		return
	}
	rnge := strings.Split(parms[0], "-")
	if len(rnge) != 2 {
		return
	}
	firstNumber, _ = strconv.Atoi(rnge[0])
	secondNumber, _ = strconv.Atoi(rnge[1])
	letter = parms[1]
	word = strings.TrimSpace(s[1])
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1NumberOfValid := 0
	part2NumberOfValid := 0
	// Size of file N
	// 1-3 c: jjjjjjjjjjjjjjjjjjjjjj
	for scanner.Scan() {
		line := scanner.Text()
		if part1(line) {
			part1NumberOfValid += 1
		}

		if part2(line) {
			part2NumberOfValid += 1
		}
	}

	fmt.Printf("Part 1: %v\n", part1NumberOfValid)
	fmt.Printf("Part 2: %v\n", part2NumberOfValid)
}
