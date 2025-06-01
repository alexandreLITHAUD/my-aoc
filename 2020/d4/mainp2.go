//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
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

	authorizedEyeColor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	if val, err := strconv.Atoi(passport.byr); err != nil || !(val >= 1920 && val <= 2002) {
		return false
	}
	if val, err := strconv.Atoi(passport.iyr); err != nil || !(val >= 2010 && val <= 2020) {
		return false
	}
	if val, err := strconv.Atoi(passport.eyr); err != nil || !(val >= 2020 && val <= 2030) {
		return false
	}

	if strings.HasSuffix(passport.hgt, "cm") {
		val, err := strconv.Atoi(passport.hgt[:len(passport.hgt)-2])
		if err != nil || val < 150 || val > 193 {
			return false
		}
	} else if strings.HasSuffix(passport.hgt, "in") {
		val, err := strconv.Atoi(passport.hgt[:len(passport.hgt)-2])
		if err != nil || val < 59 || val > 76 {
			return false
		}
	} else {
		return false
	}

	if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(passport.hcl) {
		return false
	}

	if !slices.Contains(authorizedEyeColor, passport.ecl) {
		return false
	}

	if !regexp.MustCompile(`^\d{9}$`).MatchString(passport.pid) {
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
