//go:build part2

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
	var xpos int = 0

	for i := uint(0); i < uint(len(m)); i += yDir {
		if m[i][xpos%len(m[0])] == "#" {
			trees++
		}
		xpos += int(xDir)
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

	res1 := countTrees(m, 1, 1)
	res2 := countTrees(m, 1, 3)
	res3 := countTrees(m, 1, 5)
	res4 := countTrees(m, 1, 7)
	res5 := countTrees(m, 2, 1)

	res := res1 * res2 * res3 * res4 * res5
	fmt.Printf("res1 %d, res2 %d, res3 %d, res4 %d, res5 %d\n", res1, res2, res3, res4, res5)
	fmt.Printf("Number of trees: %d\n", res)

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
