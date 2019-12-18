package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./problems.csv")
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
