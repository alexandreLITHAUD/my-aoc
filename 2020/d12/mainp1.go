//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Ship struct {
	x         int64
	y         int64
	direction rune
}

func newShip() Ship {
	return Ship{
		x:         0,
		y:         0,
		direction: 'E',
	}
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func (s *Ship) executeInstruction(instruction string) {

	runeInstruction := rune(instruction[0])
	value, err := strconv.ParseInt(instruction[1:], 10, 64)
	if err != nil {
		panic(err)
	}

	switch runeInstruction {
	case 'N':
		s.y += value
	case 'S':
		s.y -= value
	case 'E':
		s.x += value
	case 'W':
		s.x -= value
	case 'F':
		if s.direction == 'N' {
			s.y += value
		} else if s.direction == 'S' {
			s.y -= value
		} else if s.direction == 'E' {
			s.x += value
		} else if s.direction == 'W' {
			s.x -= value
		}
	case 'L':
		turns := (value / 90) % 4
		for i := int64(0); i < turns; i++ {
			s.rotateLeft90()
		}
	case 'R':
		turns := (value / 90) % 4
		for i := int64(0); i < turns; i++ {
			s.rotateRight90()
		}
	}
}

func (s *Ship) rotateLeft90() {
	switch s.direction {
	case 'N':
		s.direction = 'W'
	case 'W':
		s.direction = 'S'
	case 'S':
		s.direction = 'E'
	case 'E':
		s.direction = 'N'
	}
}

func (s *Ship) rotateRight90() {
	switch s.direction {
	case 'N':
		s.direction = 'E'
	case 'E':
		s.direction = 'S'
	case 'S':
		s.direction = 'W'
	case 'W':
		s.direction = 'N'
	}
}

func (s *Ship) parseDirectionInstruction(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		s.executeInstruction(line)
	}
	return nil
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	ship := newShip()
	err := ship.parseDirectionInstruction(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Manhattan distance: %d\n", AbsInt64(ship.x)+AbsInt64(ship.y))

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
