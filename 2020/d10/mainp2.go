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
type Memory map[uint16]int

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

	// Add outlet (0 jolts)
	adapters = append(adapters, 0)

	// Sort first to find the maximum
	slices.Sort(adapters)

	// Add device's built-in adapter (max + 3)
	maxJoltage := adapters[len(adapters)-1]
	adapters = append(adapters, maxJoltage+3)

	// Sort again to ensure correct order
	slices.Sort(adapters)

	return adapters, nil
}

func findCombination(adapters Adapters, index int, memory Memory) int {
	if index == len(adapters)-1 {
		return 1
	}

	if val, ok := memory[uint16(index)]; ok {
		return val
	}

	count := 0
	for i := index + 1; i < len(adapters); i++ {
		if adapters[i] <= adapters[index]+3 {
			count += findCombination(adapters, i, memory)
		} else {
			break
		}
	}

	memory[uint16(index)] = count
	return count
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

	mem := make(Memory)
	numberOfCombinaison := findCombination(adapters, 0, mem)
	fmt.Printf("Number of combination %d\n", numberOfCombinaison)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
