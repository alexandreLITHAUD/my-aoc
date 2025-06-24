//go:build part2

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bus uint16
type Buses []Bus
type Timestamp uint64

func parseFile(inputFile string) (Buses, Timestamp) {

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buses Buses
	var timestamp Timestamp
	scanner.Scan()
	line := scanner.Text()

	timestemp, _ := strconv.ParseUint(line, 10, 64)
	timestamp = Timestamp(timestemp)

	scanner.Scan()
	line = scanner.Text()
	strs := strings.Split(line, ",")

	for _, str := range strs {
		if str != "x" {
			bus, _ := strconv.ParseUint(str, 10, 16)
			buses = append(buses, Bus(bus))
		}
	}
	return buses, timestamp
}

func findEarliestTimestamp(buses Buses, timestamp Timestamp) uint64 {

	minValue := ^uint64(0)
	currentBus := uint64(0)
	for _, bus := range buses {
		value := uint64(math.Ceil(float64(timestamp) / float64(bus)))
		if value*uint64(bus) < minValue {
			currentBus = uint64(bus)
			minValue = value * uint64(bus)
		}
	}
	return (minValue - uint64(timestamp)) * currentBus

}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	buses, timestamp := parseFile(os.Args[1])

	fmt.Printf("Timestamp: %d\n", timestamp)

	fmt.Printf("Earliest timestamp: %d\n", findEarliestTimestamp(buses, timestamp))

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
