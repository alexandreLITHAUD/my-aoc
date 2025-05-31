//go:build part2multi

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Map [][]string

type slope struct {
	y, x uint
}

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

	slopes := []slope{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	var wg sync.WaitGroup
	results := make(chan int, len(slopes)) // buffered channel to avoid deadlock

	for _, s := range slopes {
		wg.Add(1)
		go func(s slope) {
			defer wg.Done()
			count := countTrees(m, s.y, s.x)
			results <- count
		}(s)
	}

	wg.Wait()
	close(results)

	finalResult := 1
	i := 1
	for res := range results {
		fmt.Printf("res%d = %d\n", i, res)
		finalResult *= res
		i++
	}

	elapsed := time.Since(now)

	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
