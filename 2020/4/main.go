package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filename = "input.txt"

func readFile() []string {
	//content, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//}
	//return strings.Split(string(content), "\n")

	file, err := os.Open(filename)
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

	return lines
}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *Passport) isValid() bool {
	if p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != "" {
		return true
	}
	return false
}

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
//
//    If cm, the number must be at least 150 and at most 193.
//    If in, the number must be at least 59 and at most 76.
//
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.

func validEyeColor(input string) bool {
	validColors := map[string]string{
		"amb": "amb",
		"blu": "blu",
		"brn": "brn",
		"gry": "gry",
		"grn": "grn",
		"hzl": "hzl",
		"oth": "oth",
	}

	if _, ok := validColors[input]; ok {
		return true
	}

	return false
}

func validYearRange(input string, lRange, uRange int) bool {
	if len(input) != 4 {
		return false
	}

	nInput, _ := strconv.Atoi(input)

	if nInput < lRange || nInput > uRange {
		return false
	}

	return true
}

func validHeight(input string) bool {
	measurement := input[len(input)-2:]
	if measurement == "cm" {
		value, _ := strconv.Atoi(input[:len(input)-2])
		if value >= 150 && value <= 193 {
			return true
		}
	}

	if measurement == "in" {
		value, _ := strconv.Atoi(input[:len(input)-2])
		if value >= 59 && value <= 76 {
			return true
		}
	}

	return false
}

func validHairColor(input string) bool {
	if len(input) != 7 {
		return false
	}
	match, _ := regexp.MatchString("#[a-f0-9]*", input)
	fmt.Printf("Valid Hair Color: %v\n", match)
	return match
}

func validPid(input string) bool {
	if len(input) != 9 {
		return false
	}
	match, _ := regexp.MatchString("[0-9]*", input)
	fmt.Printf("Valid PID: %v\n", match)
	return match
}

func createPassports(lines []string) []Passport {
	passports := []Passport{}

	index := 0
	for index < len(lines)-1 {
		passport := Passport{}
		fmt.Printf("Empty Passport: %+v\n", passport)
		line := lines[index]
		for line != "" {
			fmt.Printf("Index: %v\n", index)
			fmt.Printf("Line: %s\n", line)
			values := strings.Split(line, " ")
			for _, value := range values {
				if value == "" {
					continue
				}
				keyPair := strings.Split(value, ":")
				if len(keyPair) != 2 {
					continue
				}
				key := keyPair[0]
				value := keyPair[1]
				fmt.Printf("Key: %s\n", key)
				fmt.Printf("Value: %s\n", value)
				//    byr (Birth Year)
				//    iyr (Issue Year)
				//    eyr (Expiration Year)
				//    hgt (Height)
				//    hcl (Hair Color)
				//    ecl (Eye Color)
				//    pid (Passport ID)
				//    cid (Country ID)
				//byr (Birth Year) - four digits; at least 1920 and at most 2002.
				//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
				//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
				if key == "byr" && validYearRange(value, 1920, 2002) {
					passport.byr = value
				}
				if key == "iyr" && validYearRange(value, 2010, 2020) {
					passport.iyr = value
				}
				if key == "eyr" && validYearRange(value, 2020, 2030) {
					passport.eyr = value
				}
				if key == "hgt" && validHeight(value) {
					passport.hgt = value
				}
				if key == "hcl" && validHairColor(value) {
					passport.hcl = value
				}
				if key == "ecl" && validEyeColor(value) {
					passport.ecl = value
				}
				if key == "pid" && validPid(value) {
					passport.pid = value
				}
				if key == "cid" {
					passport.cid = value
				}
			}
			fmt.Printf("Current Passport: %+v\n", passport)
			index++
			if index >= len(lines) {
				break
			}

			line = lines[index]
		}

		fmt.Printf("Final Passport: %+v\n", passport)
		fmt.Printf("IsValid: %+v\n", passport.isValid())
		passports = append(passports, passport)
		index++
	}
	return passports
}

func main() {
	valid := 0
	for _, passport := range createPassports(readFile()) {
		if passport.isValid() {
			valid++
		}
	}

	fmt.Println("Result: ", valid)
}
