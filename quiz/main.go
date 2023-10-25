package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()

	lines := readFile(fileName)

	problems := ParseLines(lines)

	correct := startQuiz(problems)

	fmt.Printf("You got %d out of %d correct.\n", correct, len(problems))
}

func startQuiz(problems []Problem) int {
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)

		var ans string

		fmt.Scanf("%s\n", &ans)

		if ans == problem.answer {
			correct += 1
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
