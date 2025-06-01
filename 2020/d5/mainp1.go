//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type OnboardigBinaryMap struct {
	row string
	col string
}

func parseBinaryMap(filename string) ([]OnboardigBinaryMap, error) {

	var binaryMaps []OnboardigBinaryMap = make([]OnboardigBinaryMap, 0)

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
		binaryMaps = append(binaryMaps, OnboardigBinaryMap{
			row: line[:7],
			col: line[7:],
		})
	}

	return binaryMaps, nil
}

func findSeatID(seat OnboardigBinaryMap) (int, error) {

	var rowNumber int
	planeMinRow := 0
	planeMaxRow := 127
	for _, planeChar := range seat.row {
		if planeChar == 'F' {
			planeMaxRow = (planeMaxRow + planeMinRow) / 2
		} else if planeChar == 'B' {
			planeMinRow = (planeMaxRow+planeMinRow)/2 + 1
		} else {
			fmt.Printf("Rune not found")
			return 0, fmt.Errorf("Rune %c not found", planeChar)
		}
	}

	rowNumber = planeMaxRow

	var colNumber int
	planeMinCol := 0
	planeMaxCol := 7
	for _, planeChar := range seat.col {
		if planeChar == 'L' {
			planeMaxCol = (planeMaxCol + planeMinCol) / 2
		} else if planeChar == 'R' {
			planeMinCol = (planeMaxCol+planeMinCol)/2 + 1
		} else {
			fmt.Printf("Rune not found")
			return 0, fmt.Errorf("Rune %c not found", planeChar)
		}
	}

	colNumber = planeMaxCol

	return (rowNumber * 8) + colNumber, nil
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	binaryMaps, err := parseBinaryMap(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var maxSeatID int = 0
	for _, seat := range binaryMaps {
		seatID, err := findSeatID(seat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Printf("Max seat ID: %d\n", maxSeatID)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
