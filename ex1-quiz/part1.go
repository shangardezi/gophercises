package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName = flag.String("filename", "problems.csv", "Name of CSV file")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file with name: %s\n", *fileName))
	}

	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	problems := parseLines(lines)
	var correctAnswerCount = 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correctAnswerCount++
		}
	}

	msg := "Your Score: %d / %d"
	result := fmt.Sprintf(msg, correctAnswerCount, len(problems))
	fmt.Println(result)
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem {
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}