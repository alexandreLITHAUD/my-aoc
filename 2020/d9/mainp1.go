//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type XMASEncrytionWindow struct {
	previousNumbers    []int64
	windowSize         int
	currentWindowIndex int
}

func newXMASEncrytionWindow(windowSize int) XMASEncrytionWindow {
	return XMASEncrytionWindow{
		previousNumbers:    make([]int64, 0),
		windowSize:         windowSize,
		currentWindowIndex: 0,
	}
}

func (x *XMASEncrytionWindow) cypherInput(filename string) int64 {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var res int64 = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Bufffer initialization
		if len(x.previousNumbers) < x.windowSize {
			x.previousNumbers = append(x.previousNumbers, int64(lineInt))
			continue
		}

		find := false
		// Check if number is valid
		for _, num := range x.previousNumbers {
			for _, num2 := range x.previousNumbers {
				if num+num2 == int64(lineInt) {
					find = true
					continue
				}
			}
		}

		if !find {
			return int64(lineInt)
		}

		x.previousNumbers[x.currentWindowIndex] = int64(lineInt)
		x.currentWindowIndex = (x.currentWindowIndex + 1) % x.windowSize
	}
	return res
}

func main() {
	now := time.Now()

	windowSize := 25
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		windowSize, _ = strconv.Atoi(os.Args[2])
	}

	encrytionWindow := newXMASEncrytionWindow(windowSize)
	res := encrytionWindow.cypherInput(os.Args[1])

	fmt.Printf("Res: %d\n", res)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
}
