//go:build part2

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Questionnaire struct {
	answers map[rune]struct{}
}

func parseQuestionnaire(filename string) ([]Questionnaire, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var questionnaires []Questionnaire
	var firstPersonInGroup bool = true
	var questionnaire map[rune]struct{} = make(map[rune]struct{}, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			questionnaires = append(questionnaires, Questionnaire{answers: questionnaire})
			questionnaire = nil
			firstPersonInGroup = true
			continue
		}

		currentPerson := make(map[rune]struct{})
		for _, char := range line {
			currentPerson[char] = struct{}{}
		}

		if firstPersonInGroup {
			questionnaire = currentPerson
			firstPersonInGroup = false
		} else {
			for char := range questionnaire {
				if _, ok := currentPerson[char]; !ok {
					delete(questionnaire, char)
				}
			}
		}
	}

	// Handle last group if file does not end with a blank line
	if questionnaire != nil {
		questionnaires = append(questionnaires, Questionnaire{answers: questionnaire})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return questionnaires, nil
}

func main() {
	now := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	questionnaires, err := parseQuestionnaire(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count := 0
	for _, qu := range questionnaires {
		count += len(qu.answers)
	}

	fmt.Printf("Number of unique answers: %d\n", count)

	elapsed := time.Since(now)
	fmt.Printf("Execution time: %s\n", elapsed)
}
