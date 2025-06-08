//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var WANTED_BAG_ADJECTIVE string = "shiny"
var WANTED_BAG_COLOR string = "gold"

type Bags struct {
	adjective string
	color     string
	bagCount  map[uint16]int
	id        uint16
}

func (b Bags) String() string {
	str := fmt.Sprintf("%s %s (%d)", b.adjective, b.color, b.id)

	if len(b.bagCount) == 0 {
		str += " contains no other bags."
		return str
	}

	str += " contains: "
	parts := []string{}
	for id, count := range b.bagCount {
		parts = append(parts, fmt.Sprintf("%d x bag ID %d", count, id))
	}
	str += strings.Join(parts, ", ")
	return str
}

func getBagId(bag string, bags []Bags) (uint16, error) {
	for _, b := range bags {
		if b.adjective == strings.Fields(bag)[1] && b.color == strings.Fields(bag)[2] {
			return b.id, nil
		}
	}
	return 0, fmt.Errorf("Bag not found")
}

func parseBags(inputFile string) ([]Bags, error) {
	bags := []Bags{}

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bagsContainedString []string = make([]string, 0)
	idIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		str := strings.SplitN(line, "contain", 2)
		bagsContainedString = append(bagsContainedString, str[1])
		bags = append(bags, Bags{
			id:        uint16(idIndex),
			adjective: strings.Fields(str[0])[0],
			color:     strings.Fields(str[0])[1],
		})

		idIndex++

	}

	fmt.Printf("Number of bags: %d\n", len(bags))
	fmt.Printf("Number of bags contained: %d\n", len(bagsContainedString))

	for lines := range bagsContainedString {
		var bagContains []string = strings.Split(bagsContainedString[lines], ",")
		bags[lines].bagCount = make(map[uint16]int)
		for _, bag := range bagContains {
			bag := strings.TrimSpace(bag)
			if bag == "no other bags." {
				continue
			} else {
				bagId, err := getBagId(bag, bags)
				if err != nil {
					return nil, err
				}
				bagCount, err := strconv.Atoi(strings.Fields(bag)[0])
				if err != nil {
					return nil, err
				}
				bags[lines].bagCount[bagId] = bagCount
			}
		}
	}

	return bags, nil
}

// func holdWantedBag(bag Bags, bags []Bags, wantedBagId uint16) bool {
// 	if bag.id == wantedBagId {
// 		return true
// 	}
// 	if len(bag.bagCount) == 0 {
// 		return false
// 	}
// 	for bagId := range bag.bagCount {
// 		if holdWantedBag(bags[bagId], bags, wantedBagId) {
// 			return true
// 		}
// 	}
// 	return false
// }

func getBagIndex(bagId uint16, bags []Bags) int {
	for i, b := range bags {
		if b.id == bagId {
			return i
		}
	}
	return -1
}

func countNumberOfBags(wantedBagIndex int, bags []Bags) int {
	var counter int
	for bagId, bagCount := range bags[wantedBagIndex].bagCount {
		counter += bagCount * (1 + countNumberOfBags(getBagIndex(bagId, bags), bags))
	}
	return counter
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	bags, err := parseBags(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, b := range bags {
		fmt.Printf("%s\n", b.String())
	}

	// var wantedBagId uint16
	var wantedBagIndex int
	for i, b := range bags {
		if b.adjective == WANTED_BAG_ADJECTIVE && b.color == WANTED_BAG_COLOR {
			// wantedBagId = b.id
			wantedBagIndex = i
		}
	}

	numberOfBags := countNumberOfBags(wantedBagIndex, bags)

	fmt.Printf("Number of bags that can hold %s %s: %d\n", WANTED_BAG_ADJECTIVE, WANTED_BAG_COLOR, numberOfBags)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
