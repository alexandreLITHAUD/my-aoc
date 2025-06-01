//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

type OnboardingBinaryMap struct {
	row string
	col string
}

func parseBinaryMap(filename string) ([]OnboardingBinaryMap, error) {

	var binaryMaps []OnboardingBinaryMap = make([]OnboardingBinaryMap, 0)

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
		binaryMaps = append(binaryMaps, OnboardingBinaryMap{
			row: line[:7],
			col: line[7:],
		})
	}

	return binaryMaps, nil
}

func findSeatID(seat OnboardingBinaryMap) (int, error) {

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

	if rowNumber == 0 || rowNumber == 127 {
		return -1, nil
	}

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

	seatIDs := make([]int, 0)
	for _, seat := range binaryMaps {
		seatID, err := findSeatID(seat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if seatID == -1 {
			continue
		}
		seatIDs = append(seatIDs, seatID)
	}
	slices.Sort(seatIDs)

	var missingSeatID int
	for i := 1; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1]-seatIDs[i] == 2 {
			missingSeatID = seatIDs[i] + 1
			break
		}
	}
	fmt.Printf("My seat ID: %d\n", missingSeatID)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
