package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("filename", "problems.csv", "The CSV file's name in the format of 'question,answer'")
	flag.Parse()

	fmt.Println("Press enter to start the quiz")
	buf := bufio.NewReader(os.Stdin)
	_, err := buf.ReadBytes('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		quiz(csvFilename)
	}
}

func quiz(fileName *string) {
	data, err := ioutil.ReadFile(*fileName)

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	reader := bufio.NewReader(os.Stdin)
	score := 0
	totalQuestions := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
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
