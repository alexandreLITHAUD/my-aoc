//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Password struct {
	corruptedPassword string
	min               uint8
	max               uint8
	char              rune
}

func parseFile(filename string) ([]Password, error) {
	var passwords []Password

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var password Password
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &password.min, &password.max, &password.char, &password.corruptedPassword)
		if err != nil {
			return nil, err
		}

		passwords = append(passwords, password)
	}
	return passwords, nil
}

func isPasswordValid(password Password) bool {
	count := 0
	for _, char := range password.corruptedPassword {
		if char == password.char {
			count++
		}
	}
	return count >= int(password.min) && count <= int(password.max)
}

func main() {

	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	passwords, err := parseFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var count int
	for _, password := range passwords {
		if isPasswordValid(password) {
			count++
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", count)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
