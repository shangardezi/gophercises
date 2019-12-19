package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("filename", "problems.csv", "The CSV file's name in the format of 'question,answer'")
	timeLimit := flag.Int("timelimit", 30, "the duration of the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Unable to open file named: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Unable to read CSV file")
	}

	fmt.Println("Press enter to start the quiz")

	buf := bufio.NewReader(os.Stdin)
	_, err2 := buf.ReadBytes('\n')

	if err2 != nil {
		exit("Unable to start quiz")
	}

	problems := parseLines(lines)
	score := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem %d of %d: %s\n", i+1, len(problems), p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time's up! You scored: %d/%d\n", score, len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a {
				score++
			}
		}

		fmt.Printf("\nYou scored: %d/%d\n", score, len(problems))
	}
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return problems
}
