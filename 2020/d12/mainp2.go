//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Ship struct {
	x int64
	y int64
}

type Waypoint struct {
	x int64 // relative to ship
	y int64 // relative to ship
}

func newShip() Ship {
	return Ship{
		x: 0,
		y: 0,
	}
}

func newWaypoint() Waypoint {
	return Waypoint{
		x: 10, // 10 units east of ship
		y: 1,  // 1 unit north of ship
	}
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Rotate waypoint left by 90 degrees around the ship
func (w *Waypoint) rotateLeft90() {
	// (x, y) -> (-y, x)
	newX := -w.y
	newY := w.x
	w.x = newX
	w.y = newY
}

// Rotate waypoint right by 90 degrees around the ship
func (w *Waypoint) rotateRight90() {
	// (x, y) -> (y, -x)
	newX := w.y
	newY := -w.x
	w.x = newX
	w.y = newY
}

func (s *Ship) executeInstruction(instruction string, waypoint *Waypoint) {
	runeInstruction := rune(instruction[0])
	value, err := strconv.ParseInt(instruction[1:], 10, 64)
	if err != nil {
		panic(err)
	}

	switch runeInstruction {
	case 'N':
		waypoint.y += value
	case 'S':
		waypoint.y -= value
	case 'E':
		waypoint.x += value
	case 'W':
		waypoint.x -= value
	case 'F':
		// Move ship toward waypoint, value times
		s.x += waypoint.x * value
		s.y += waypoint.y * value
		// Waypoint maintains its relative position to ship
	case 'L':
		// Rotate waypoint left around ship
		turns := (value / 90) % 4
		for i := int64(0); i < turns; i++ {
			waypoint.rotateLeft90()
		}
	case 'R':
		// Rotate waypoint right around ship
		turns := (value / 90) % 4
		for i := int64(0); i < turns; i++ {
			waypoint.rotateRight90()
		}
	}
}

func (s *Ship) parseDirectionInstruction(filename string, waypoint *Waypoint) error {
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
		s.executeInstruction(line, waypoint)
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
	waypoint := newWaypoint()

	err := ship.parseDirectionInstruction(os.Args[1], &waypoint)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Ship final position: (%d, %d)\n", ship.x, ship.y)
	fmt.Printf("Waypoint relative position: (%d, %d)\n", waypoint.x, waypoint.y)
	fmt.Printf("Manhattan distance: %d\n", AbsInt64(ship.x)+AbsInt64(ship.y))
	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
}
