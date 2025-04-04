package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Submarine struct {
	depth int
}

func checkDeapth(filename string) uint {

	var total uint = 0
	var sub Submarine

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		deapth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sub = Submarine{
			depth: deapth,
		}
	} else {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	for scanner.Scan() {
		deapth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if sub.depth < deapth {
			total++
		}
		sub.depth = deapth
	}
	return total

}

func main() {
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	var total uint = checkDeapth(filename)

	fmt.Printf("Total: %d\n", total)
}
