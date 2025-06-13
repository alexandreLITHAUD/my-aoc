//go:build part2

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
	acc  int16
	pc   int16
}

func NewConsoleMemory() ConsoleMemory {
	var c ConsoleMemory
	c.seen = make(map[int]struct{})
	c.acc = 0
	c.pc = 0
	return c
}

func parseScriptToMemoryInstrcutions(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var scriptIns []Instruction = make([]Instruction, 0)
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
		scriptIns = append(scriptIns, ins)
	}

	return scriptIns
}

func (c *ConsoleMemory) CleanMemory() {
	c.seen = make(map[int]struct{})
	c.acc = 0
	c.pc = 0
}

func (c *ConsoleMemory) Execute(ins []Instruction) (int16, bool) {
	for {
		if c.pc >= int16(len(ins)) {
			break
		}
		if _, ok := c.seen[int(c.pc)]; ok {
			return c.acc, false
		}

		c.seen[int(c.pc)] = struct{}{}
		switch ins[c.pc].operation {
		case "nop":
			c.pc++
		case "acc":
			c.acc += int16(ins[c.pc].value)
			c.pc++
		case "jmp":
			c.pc += int16(ins[c.pc].value)
		}
	}
	return c.acc, true
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	consoleMemory := NewConsoleMemory()
	instructions := parseScriptToMemoryInstrcutions(os.Args[1])

	for i := 0; i < len(instructions); i++ {
		originalOp := instructions[i].operation
		if originalOp == "nop" {
			instructions[i].operation = "jmp"
		} else if originalOp == "jmp" {
			instructions[i].operation = "nop"
		} else {
			continue
		}

		acc, ok := consoleMemory.Execute(instructions)
		consoleMemory.CleanMemory()
		if ok {
			fmt.Printf("acc: %d\n", acc)
			break
		}

		// Restore
		instructions[i].operation = originalOp
	}

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
	os.Exit(0)
}
