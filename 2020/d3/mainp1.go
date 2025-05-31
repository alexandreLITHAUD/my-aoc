//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Map [][]string

func parseFile(inputFile string) (Map, error) {

	m := Map(make([][]string, 0))

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

func (m Map) print() {
	for _, row := range m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func countTrees(m Map, yDir uint, xDir uint) int {
	var trees int = 0

	for i := uint(0); i < uint(len(m)); i += yDir {
		xpos := (int(i) * int(xDir)) % len(m[0])
		if m[i][xpos] == "#" {
			trees++
		}
	}
	return trees
}

func main() {

	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	m, err := parseFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m.print()

	res := countTrees(m, 1, 3)

	fmt.Printf("Number of trees: %d\n", res)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
