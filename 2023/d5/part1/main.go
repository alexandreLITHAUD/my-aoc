package main

import (
	"bufio"
	"fmt"
	"os"
)

type Farm struct {
	seeds             []uint8
	fertilizerSeedMap map[uint8]uint8
}

func parseFile(inputFile string) (Farm, error) {

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		
	}

	return Farm{}, nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	fmt.Printf("Hello World")
}
