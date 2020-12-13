package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EnriqueL8/adventofcode/utils"
)

const filename = "input.txt"

func extractValues(line string) (string, string) {
	values := strings.Split(line, " ")
	if len(values) != 2 {
		return "", ""
	}

	return values[0], values[1]
}

func loop(lines []string) (int, bool) {
	cmds := map[int]string{}
	acc := 0
	i := 0
	for i < len(lines)-1 {
		cmd, value := extractValues(lines[i])
		if cmd == "" || value == "" {
			i++
			continue
		}

		if _, ok := cmds[i]; ok {
			return acc, false
		}

		cmds[i] = cmd

		if cmd == "nop" {
			i++
			continue
		}

		if cmd == "acc" {
			n, _ := strconv.Atoi(value[1:])
			if string(value[0]) == "-" {
				acc -= n
			} else if string(value[0]) == "+" {
				acc += n
			}
			i++
			continue
		}

		if cmd == "jmp" {
			n, _ := strconv.Atoi(value[1:])
			if string(value[0]) == "-" {
				i -= n
			} else if string(value[0]) == "+" {
				i += n
			}
		}
	}
	return acc, true
}

func main() {
	lines, _ := utils.ReadLines("input.txt", "\n")

	acc1, _ := loop(lines)
	fmt.Println("Part 1:", acc1)

	for i, line := range lines {
		nLines := lines
		cmd, value := extractValues(line)
		if cmd == "nop" {
			nLines[i] = fmt.Sprintf("%s %s", "jmp", value)
			acc, exit := loop(nLines)
			if exit {
				fmt.Println("Part 2:", acc)
				break
			}
			nLines[i] = fmt.Sprintf("%s %s", "nop", value)
		}

		if cmd == "jmp" {
			nLines[i] = fmt.Sprintf("%s %s", "nop", value)
			acc, exit := loop(nLines)
			if exit {
				fmt.Println("Part 2:", acc)
				break
			}
			nLines[i] = fmt.Sprintf("%s %s", "jmp", value)
		}
	}
}
