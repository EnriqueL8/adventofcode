package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/EnriqueL8/adventofcode/utils"
)

func getNumbers() (numbers []int) {
	lines, _ := utils.ReadLines("input.txt", "\n")
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	return
}

func findSum(n int, numbers []int) bool {
	sNumbers := make([]int, len(numbers))
	copy(sNumbers, numbers)
	sort.Ints(sNumbers)
	i := 0
	v := len(sNumbers) - 1
	for i != v {
		sum := sNumbers[i] + sNumbers[v]
		if sum == n {
			return true
		}
		if sum > n {
			v--
		}

		if sum < n {
			i++
		}
	}

	return false
}

func sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func findSequence(n int, numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for v := i + 1; v < len(numbers)-1; v++ {
			if n == sum(numbers[i:v]) {
				return numbers[i:v]
			}
		}
	}
	return []int{}
}

func main() {
	numbers := getNumbers()
	preamble := 25
	i := preamble
	invalid := 0
	for i < len(numbers)-1 {
		found := findSum(numbers[i], numbers[i-preamble:i])
		if found {
			i++
			continue
		}

		invalid = numbers[i]
		break
	}
	fmt.Println("Part 1: ", invalid)

	seq := findSequence(invalid, numbers)
	sort.Ints(seq)
	if len(seq) > 1 {
		weakness := seq[0] + seq[len(seq)-1]
		fmt.Println("Part 2: ", weakness)
	}

}
