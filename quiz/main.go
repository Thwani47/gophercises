package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "An indicator of whether the problem list should be shuffled or not")

	flag.Parse()

	lines := readFile(fileName)

	problems := ParseLines(lines)

	if *shuffle == true {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	correct := startQuiz(problems, *timeLimit)

	fmt.Printf("\nYou got %d out of %d correct.\n", correct, len(problems))
}

func startQuiz(problems []Problem, quizDuration int) int {
	timer := time.NewTimer(time.Duration(*&quizDuration) * time.Second)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		answerChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChan <- ans
		}()
		select {
		case <-timer.C:
			return correct
		case answer := <-answerChan:
			if answer == problem.answer {
				correct += 1
			}
		}
	}

	return correct
}

func readFile(fileName *string) [][]string {
	file, err := os.Open(*fileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *fileName))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
