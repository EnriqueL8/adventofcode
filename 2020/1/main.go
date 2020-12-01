package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	// Brute force
	for _, v := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if v+j+k == 2020 {
					multi := v * j * k
					fmt.Printf("Result: %v\n", multi)
					os.Exit(0)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
