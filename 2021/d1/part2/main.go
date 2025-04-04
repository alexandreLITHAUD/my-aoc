package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Windows struct {
	depth int
	size  uint
}

func checkWindows(windows []*Windows, windowSize uint) (uint, error) {

	var total uint = 0

	for i := range windows {

		if int(windowSize) != int(windows[i].size) {
			continue
		}

		if i+1 >= len(windows) {
			break
		}

		if windows[i].depth < windows[i+1].depth {
			total++
		}

	}

	return total, nil
}

func checkDeapthWindow(filename string, windowSize uint) (uint, error) {

	var windows []*Windows

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		deapth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, win := range windows {
			if int(win.size) != int(windowSize) {
				win.depth += deapth
				win.size++
			}
		}

		windows = append(windows, &Windows{
			depth: deapth,
			size:  1,
		})

	}

	return checkWindows(windows, windowSize)

}

func main() {
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	total, err := checkDeapthWindow(filename, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Total: %d\n", total)
}
