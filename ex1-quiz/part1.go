package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("filename", "problems.csv", "The CSV file's name in the format of 'question,answer'")
	flag.Parse()

	data, err := ioutil.ReadFile(*csvFilename)

	if err != nil {
		error(fmt.Sprintf("Unable to open file named: %s", *csvFilename))
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	fmt.Println("Press enter to start the quiz")

	buf := bufio.NewReader(os.Stdin)
	_, err2 := buf.ReadBytes('\n')

	if err2 != nil {
        error("Unable to start quiz")
	} else {
		quiz(r)
	}
}

func error(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func quiz(r *csv.Reader) {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	totalQuestions := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
            error("Error reading file data")
		}

		fmt.Println(record[0])
		text, _ := reader.ReadString('\n')

		if strings.TrimRight(text, "\n") == record[1] {
			score++
		}

		totalQuestions++
	}

	fmt.Println("You scored:", score, "/", totalQuestions)
}
