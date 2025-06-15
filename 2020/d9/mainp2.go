//go:build part2

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

type DecryptionWindow struct {
	window         []int64
	minWindowValue int64
	maxWindowValue int64
	total          int64
}

func newXMASEncrytionWindow(windowSize int) XMASEncrytionWindow {
	return XMASEncrytionWindow{
		previousNumbers:    make([]int64, 0),
		windowSize:         windowSize,
		currentWindowIndex: 0,
	}
}

func newDecryptionWindow() DecryptionWindow {
	return DecryptionWindow{
		window:         make([]int64, 0),
		minWindowValue: int64(^uint64(0) >> 1),
		maxWindowValue: 0,
		total:          0,
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

		lineInt, err := strconv.ParseInt(line, 10, 64)
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
		for i, num := range x.previousNumbers {
			for j, num2 := range x.previousNumbers {
				if i != j && num+num2 == int64(lineInt) {
					find = true
					break
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

func (x *XMASEncrytionWindow) decypherXMAS(cyperToken int64, filename string) DecryptionWindow {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var DecryptionWindows []DecryptionWindow = make([]DecryptionWindow, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		lineInt, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create DecryptionWindow
		dw := newDecryptionWindow()
		DecryptionWindows = append(DecryptionWindows, dw)

		newWindows := make([]DecryptionWindow, 0, len(DecryptionWindows))
		for _, elem := range DecryptionWindows {

			elem.window = append(elem.window, int64(lineInt))
			elem.total += int64(lineInt)

			// update Max and Min
			if elem.maxWindowValue < int64(lineInt) {
				elem.maxWindowValue = int64(lineInt)
			}
			if elem.minWindowValue > int64(lineInt) {
				elem.minWindowValue = int64(lineInt)
			}

			if elem.total < cyperToken {
				newWindows = append(newWindows, elem)
			}

			if elem.total == cyperToken {
				return elem
			}

		}
		DecryptionWindows = newWindows

	}
	return DecryptionWindow{}

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
	cyperToken := encrytionWindow.cypherInput(os.Args[1])

	cleanEncryptionWindow := newXMASEncrytionWindow(windowSize)
	weaknessToken := cleanEncryptionWindow.decypherXMAS(cyperToken, os.Args[1])

	fmt.Printf("Encryption Weakness : %d\n", weaknessToken.minWindowValue+weaknessToken.maxWindowValue)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
}
