//go:build part1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Ship struct {
	x int64
	y int64
}

type Waypoint struct {
	x int64
	y int64
}

func calculateDirectionVector(ship Ship, waypoint Waypoint) (int64, int64) {
	return waypoint.x - ship.x, waypoint.y - ship.y
}

func newShip() Ship {
	return Ship{
		x: 0,
		y: 0,
	}
}

func newWaypoint() Waypoint {
	return Waypoint{
		x: 10,
		y: 1,
	}
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func (w *Waypoint) calculatePositionRotation90(ship Ship, direction int) {
	angle := float64(direction) * 90.0 * math.Pi / 180.0
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	xpos := w.x - ship.x
	ypos := w.y - ship.y

	w.x = int64(float64(xpos)*cos-float64(ypos)*sin) + ship.x
	w.y = int64(float64(xpos)*sin+float64(ypos)*cos) + ship.y
}

func (s *Ship) executeInstruction(instruction string) {

	runeInstruction := rune(instruction[0])
	value, err := strconv.ParseInt(instruction[1:], 10, 64)
	if err != nil {
		panic(err)
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
