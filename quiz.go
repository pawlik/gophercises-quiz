package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, er := os.Open("problems.csv")
	logAndExitIfError(er)

	reader := csv.NewReader(file)

	records, er := reader.ReadAll()
	logAndExitIfError(er)
	correctAnswers := 0
	current := 0
	for _, record := range records {
		question, answer := record[0], record[1]
		fmt.Print(question, "= ")
		correctAnswers += checkAnswer(answer)
		current++
	}

	fmt.Fprintf(os.Stdout, "Correct %v out of %v\n", correctAnswers, current)
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
