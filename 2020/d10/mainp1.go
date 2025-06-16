//go:build part1

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

	adapters = append(adapters, 0)
	slices.Sort(adapters)
	return adapters, nil
}

func joltDifference(adapter Adapters) int {
	onediff := 0
	threediff := 0

	for i := 1; i < len(adapter); i++ {
		diff := int(adapter[i] - adapter[i-1])
		if diff == 1 {
			onediff++
		} else if diff == 3 {
			threediff++
		}
	}

	threediff++
	fmt.Printf("Number of 1 diff %d, Number if 3 diff %d\n", onediff, threediff)

	return onediff * threediff
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

	diff := joltDifference(adapters)

	fmt.Printf("Total differences: %d\n", diff)
	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
