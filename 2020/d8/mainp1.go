//go:build part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Instruction struct {
	operation string
	value     int
}

type ConsoleMemory struct {
	seen map[int]struct{}
	ins  []Instruction
	acc  int16
	pc   int16
}

func NewConsoleMemory() ConsoleMemory {
	var c ConsoleMemory
	c.seen = make(map[int]struct{})
	c.ins = make([]Instruction, 0)
	c.acc = 0
	c.pc = 0
	return c
}

func (c *ConsoleMemory) parseScriptToConsoleMemory(filename string) {
	file, err := os.Open(filename)
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

		var ins Instruction
		_, err := fmt.Sscanf(line, "%s %d", &ins.operation, &ins.value)
		if err != nil {
			panic(err)
		}
		c.ins = append(c.ins, ins)
	}
}

func (c *ConsoleMemory) Execute() int16 {
	for {
		if c.pc >= int16(len(c.ins)) {
			break
		}
		if _, ok := c.seen[int(c.pc)]; ok {
			return c.acc
		}

		c.seen[int(c.pc)] = struct{}{}
		switch c.ins[c.pc].operation {
		case "nop":
			c.pc++
		case "acc":
			c.acc += int16(c.ins[c.pc].value)
			c.pc++
		case "jmp":
			c.pc += int16(c.ins[c.pc].value)
		}
	}
	return -1
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	consoleMemory := NewConsoleMemory()
	consoleMemory.parseScriptToConsoleMemory(os.Args[1])
	acc := consoleMemory.Execute()
	fmt.Printf("acc: %d\n", acc)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
