package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	userReader := bufio.NewReader(os.Stdin)
	var lineCount = 0
	var correctAnswerCount = 0
	for {
		line, error := reader.Read()
			if error == io.EOF {
				break //End of file, stop looping
		} else if error != nil {
			log.Fatal(error) // Something went wrong, log fatal error
			}
		lineCount += 1
		fmt.Println(line[0])
		text, _ := userReader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == line[1] {
			fmt.Println("Correct!")
			correctAnswerCount += 1
		} else {
			fmt.Println("Uh oh, that's not right")
		}
	}

	msg := "Your Score: %d / %d"
	result := fmt.Sprintf(msg, correctAnswerCount, lineCount)
	fmt.Println(result)
}
