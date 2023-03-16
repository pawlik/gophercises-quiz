package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	records := getQuestions()

	go quizLoop(records)

	<-timer.C

	// fmt.Fprintf(os.Stdout, "Correct %v out of %v\n", correctAnswers, len(records))
}

func quizLoop(records [][]string) {
	for _, record := range records {
		question, answer := record[0], record[1]
		fmt.Print(question, "= ")
		checkAnswer(answer)
	}
}

func getQuestions() [][]string {
	file, er := os.Open("problems.csv")
	logAndExitIfError(er)

	reader := csv.NewReader(file)

	records, er := reader.ReadAll()
	logAndExitIfError(er)
	return records
}

// Returns 1 when correct answer, 0 otherwise.
func checkAnswer(answer string) int {
	reader := bufio.NewReader(os.Stdin)
	providedAnswer, er := reader.ReadString('\n')
	logAndExitIfError(er)
	if strings.Trim(providedAnswer, "\n") == answer {
		return 1
	}
	return 0
}

func logAndExitIfError(er error) {
	if er != nil {
		fmt.Println("Error", er)
		os.Exit(1)
	}
}
