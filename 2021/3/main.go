package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseNumbers(numbersInBits []string) ([]int, int) {
	nOfBits := len(numbersInBits[0]) - 1
	numbers := []int{}
	for _, n := range numbersInBits {
		numbers = append(numbers, bitToInt(n))
	}
	return numbers, nOfBits
}

func part1(numbers []int, nOfBits int) int {
	gamma := 0
	epsilon := 0

	half := len(numbers) / 2
	for bit := (1 << nOfBits); bit >= 1; bit /= 2 {
		counter := 0
		for _, n := range numbers {
			// Keep n here!
			if (n & bit) != 0 {
				counter++
			}
		}
		if counter >= half {
			gamma += bit
		} else {
			epsilon += bit
		}
	}

	return gamma * epsilon
}

/*
  pass true for most bits
  pass false for less bits
*/
func reduce(numbers []int, nOfBits int, filter bool) int {
	mask := 0
	val := 0
	for bit := (1 << nOfBits); bit >= 1; bit /= 2 {
		counter := 0
		total := 0
		for _, n := range numbers {
			// Current number does not match the mask up till now
			if n&mask != val {
				continue
			}
			total++
			if n&bit != 0 {
				counter++
			}
		}
		if total == 1 {
			// You have to return early here because in the false case
			// it will always be more than half so will never increase the val
			return match(numbers, val, mask)
		}

		moreThanHalf := counter >= (total+1)/2
		if moreThanHalf == filter {
			val += bit
		}
		mask += bit
	}
	return val
}

func match(numbers []int, val, mask int) int {
	for _, n := range numbers {
		if n&mask == val {
			return n
		}
	}
	return -1
}

func part2(numbers []int, nOfBits int) int {
	return reduce(numbers, nOfBits, true) * reduce(numbers, nOfBits, false)
}

func bitToInt(bit string) int {
	i, err := strconv.ParseInt(bit, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(i)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbersInBits := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbersInBits = append(numbersInBits, scanner.Text())
	}

	numbres, nOfBits := parseNumbers(numbersInBits)

	fmt.Printf("Part1: %v\n", part1(numbres, nOfBits))
	fmt.Printf("Part2: %v\n", part2(numbres, nOfBits))
}
