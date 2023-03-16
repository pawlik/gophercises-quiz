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
	timer := time.NewTimer(360 * time.Second)
	records := getQuestions()

	answersChannel := make(chan int)
	done := make(chan bool)

	go quizLoop(records, answersChannel, done)

	total := 0

	for {
		select {
		case a := <-answersChannel:
			total += a
		case <-timer.C:
			fmt.Println(
				"Times up!",
				fmt.Sprintf("\nYour score: %v/%v", total, len(records)),
			)
			return
		case <-done:
			fmt.Println(
				"Nice, you answered all questions!",
				fmt.Sprintf("\nYour score: %v/%v", total, len(records)),
			)
			return
		}
	}
}

func quizLoop(records [][]string, answersChannel chan int, done chan bool) {
	for _, record := range records {
		question, answer := record[0], record[1]
		fmt.Print(question, "= ")
		answersChannel <- checkAnswer(answer)
	}
	done <- true
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
