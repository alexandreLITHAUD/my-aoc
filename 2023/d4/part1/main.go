package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winninNumbers []uint16
	playedNumbers []uint16
}

func (c Card) cardPoint() uint16 {
	var sum uint16 = 0

	for _, num := range c.winninNumbers {
		if slices.Contains(c.playedNumbers, num) {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}
	return sum
}

func parseNumbers(numbers string) []uint16 {
	var result []uint16
	for _, num := range strings.Split(numbers, " ") {
		if num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			result = append(result, uint16(n))
		}
	}
	return result
}

func parseFile(inputFile string) ([]Card, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cards []Card
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		re := regexp.MustCompile(`Card\s+\d+:\s+([\d\s]+)\|([\d\s]+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			cards = append(cards, Card{
				winninNumbers: parseNumbers(matches[1]),
				playedNumbers: parseNumbers(matches[2]),
			})
		} else {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
	}
	return cards, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	cards, err := parseFile(inputFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sum uint16
	for _, card := range cards {
		sum += card.cardPoint()
	}

	fmt.Printf("Sum of card points: %d\n", sum)
}
