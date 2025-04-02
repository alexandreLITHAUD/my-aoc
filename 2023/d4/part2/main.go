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
	instance      uint16
}

func UpdateInstance(index uint16, cards map[uint16]*Card) {

	c := cards[index]

	var somme uint8 = 0
	for _, num := range c.winninNumbers {
		if slices.Contains(c.playedNumbers, num) {
			somme++
		}
	}

	for i := uint8(0); i < somme; i++ {
		newIndex := index + 1 + uint16(i)
		if nextCard, exists := cards[newIndex]; exists {
			nextCard.instance += c.instance
		}
	}

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

func parseFile(inputFile string) (map[uint16]*Card, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cards map[uint16]*Card = make(map[uint16]*Card)
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
		re := regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+)\|([\d\s]+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {

			var id uint16 = parseNumbers(matches[1])[0]
			var card *Card = &Card{
				winninNumbers: parseNumbers(matches[2]),
				playedNumbers: parseNumbers(matches[3]),
				instance:      1,
			}
			cards[id] = card

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

	for i, _ := range cards {
		UpdateInstance(i+1, cards)
	}

	var sum uint16
	for _, card := range cards {
		sum += card.instance
	}

	fmt.Printf("Sum of card instances: %d\n", sum)
}
