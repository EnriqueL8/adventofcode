package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Toboggan struct {
	lines  []string
	length int
	width  int
}

func (t *Toboggan) calculateSlope(sx, sy int) int {
	x := 0
	y := 0
	trees := 0
	for y < t.length-1 {
		x += sx
		y += sy
		if t.lines[y][x%t.width] == '#' {
			trees += 1
		}
	}

	return trees
}

func (t *Toboggan) part1() int {
	return t.calculateSlope(3, 1)
}

func (t *Toboggan) part2() int {
	return t.calculateSlope(1, 1) * t.calculateSlope(3, 1) * t.calculateSlope(5, 1) * t.calculateSlope(7, 1) * t.calculateSlope(1, 2)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	t := &Toboggan{lines, len(lines), len(lines[0])}

	fmt.Printf("Part1 : %v\n", t.part1())
	fmt.Printf("Part2 : %v\n", t.part2())
}
