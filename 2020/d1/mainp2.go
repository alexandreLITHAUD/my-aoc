//go:build part2

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

	for i := 0; i < len(res.Numbers); i++ {
		for j := i + 1; j < len(res.Numbers); j++ {
			for k := j + 1; k < len(res.Numbers); k++ {
				if res.Numbers[i]+res.Numbers[j]+res.Numbers[k] == goal {
					fmt.Printf("Found numbers %d, %d and %d\n Res is %d\n", res.Numbers[i], res.Numbers[j], res.Numbers[k], res.Numbers[i]*res.Numbers[j]*res.Numbers[k])
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println("No match found")
	os.Exit(1)
}
