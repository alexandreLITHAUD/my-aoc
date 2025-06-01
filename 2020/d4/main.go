//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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

func parsePassport(filename string) ([]Passport, error) {

	var passports []Passport = make([]Passport, 0)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var passport Passport

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, passport)
			passport = Passport{}
		}

		elemInCurLine := strings.Fields(line)
		for _, elem := range elemInCurLine {
			var field string = ""
			var value string = ""

			parts := strings.SplitN(elem, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid field format: %q", elem)
			}
			field = parts[0]
			value = parts[1]

			switch field {
			case "byr":
				passport.byr = value
			case "iyr":
				passport.iyr = value
			case "eyr":
				passport.eyr = value
			case "hgt":
				passport.hgt = value
			case "hcl":
				passport.hcl = value
			case "ecl":
				passport.ecl = value
			case "pid":
				passport.pid = value
			case "cid":
				passport.cid = value
			}
		}
	}

	if passport != (Passport{}) {
		passports = append(passports, passport)
	}

	return passports, nil
}

func checkPassportValidity(passport Passport) bool {
	if passport.byr == "" ||
		passport.iyr == "" ||
		passport.eyr == "" ||
		passport.hgt == "" ||
		passport.hcl == "" ||
		passport.ecl == "" ||
		passport.pid == "" {
		return false
	}
	return true
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	res, err := parsePassport(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var counter int = 0
	for _, passport := range res {
		if checkPassportValidity(passport) {
			counter++
		}
	}

	fmt.Printf("Number of valid passports: %d\n", counter)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
