//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Adapter uint16
type Adapters []Adapter

func parseAllAdapters(filename string) (Adapters, error) {
	adapters := make(Adapters, 0)

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

		joltage, err := strconv.ParseUint(line, 10, 16)
		if err != nil {
			return nil, err
		}
		adapters = append(adapters, Adapter(joltage))
	}

	adapters = append(adapters, adapters[len(adapters)-1]+3)
	adapters = append(adapters, 0)
	slices.Sort(adapters)
	return adapters, nil
}

func findCombination(adapters Adapters) int {

	return 0
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	adapters, err := parseAllAdapters(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numberOfCombinaison := findCombination(adapters)
	fmt.Printf("Number of combination %d\n", numberOfCombinaison)
	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
