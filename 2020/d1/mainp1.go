//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FileNumbers struct {
	Numbers []uint64
}

func parseFile(filename string) (*FileNumbers, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numArr := FileNumbers{
		Numbers: []uint64{},
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		numArr.Numbers = append(numArr.Numbers, uint64(num))
	}
	return &numArr, nil
}

func findSum(a uint64, b uint64, goal uint64) bool {
	return a+b == goal
}

func main() {
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	var goal uint64 = 2020
	if len(os.Args) > 2 {
		tmpgoal, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Errorf(err.Error())
			os.Exit(1)
		}
		goal = uint64(tmpgoal)
	}

	filename := os.Args[1]
	res, err := parseFile(filename)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	seen := make(map[uint64]struct{})

	for _, num := range res.Numbers {
		expected := goal - num
		if _, ok := seen[expected]; ok {
			fmt.Printf("Found numbers %d and %d\n Res is %d\n", num, expected, num*expected)
			os.Exit(0)
		}
		seen[num] = struct{}{}
	}

	fmt.Println("No match found")
	os.Exit(1)
}
