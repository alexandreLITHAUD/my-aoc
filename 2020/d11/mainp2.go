//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Train [][]string

func parseTrainState(inputFile string) (Train, error) {

	m := Train(make([][]string, 0))

	file, err := os.Open(inputFile)
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
		m = append(m, strings.Split(line, ""))
	}
	return m, nil
}

func (t Train) printTrainState() {
	for _, row := range t {
		fmt.Println(row)
	}
	fmt.Println()
}

func (t Train) getCell(row int, col int) rune {

	if row < 0 || row >= len(t) {
		return '.'
	}
	if col < 0 || col >= len(t[0]) {
		return '.'
	}

	return rune(t[row][col][0])
}

func (t *Train) setCellOccupied(row int, col int) {
	(*t)[row][col] = "#"
}

func (t *Train) setCellEmpty(row int, col int) {
	(*t)[row][col] = "L"
}

func (t Train) getNeighborsOccupiedNumber(row int, col int) int {

	neighbors := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, neighbor := range neighbors {
		xcoord := row
		ycoord := col

		for {

			xcoord += neighbor[0]
			ycoord += neighbor[1]

			if xcoord < 0 || xcoord >= len(t) {
				break
			}
			if ycoord < 0 || ycoord >= len(t[0]) {
				break
			}
			if t.getCell(xcoord, ycoord) == '#' {
				count++
				break
			} else if t.getCell(xcoord, ycoord) == 'L' {
				break
			}
		}
	}

	return count
}

func (t Train) countNumberOfOccupiedSeats() int {
	count := 0
	for _, row := range t {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}
	return count
}

func (t *Train) simulateTrainStateRound() bool {

	otherTrain := Train(make([][]string, len(*t)))
	for i := range *t {
		otherTrain[i] = make([]string, len((*t)[i]))
		copy(otherTrain[i], (*t)[i])
	}

	changed := false
	for i, row := range *t {
		for j, col := range row {
			if col == "L" {
				if t.getNeighborsOccupiedNumber(i, j) == 0 {
					otherTrain.setCellOccupied(i, j)
					changed = true
				}
			} else if col == "#" {
				if t.getNeighborsOccupiedNumber(i, j) >= 5 {
					otherTrain.setCellEmpty(i, j)
					changed = true
				}
			}
		}
	}

	*t = otherTrain
	return changed
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	train, err := parseTrainState(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	train.printTrainState()
	for train.simulateTrainStateRound() {
		continue
	}
	train.printTrainState()
	fmt.Printf("Number of occupied seats: %d\n", train.countNumberOfOccupiedSeats())

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
