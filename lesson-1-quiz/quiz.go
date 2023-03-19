package main

// go build quiz.go && ./quiz
import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const DEFAULT_FILE string = "problems.csv"
const DEFAULT_TIME_SECONDS int = 30

func main() {
	fileNameFlag := flag.String("csv", DEFAULT_FILE, "path to file you want to use")
	// timeLimitFlag := flag.Int("limit", DEFAULT_TIME_SECONDS, "time limit per question")
	flag.Parse()

	f, err := os.Open(*fileNameFlag)

	if err != nil {
		log.Fatalf("We could not open the CSV file: %s", *fileNameFlag)
	}

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Failed to parse the CSV file.")
	}

	if records == nil {
		log.Fatal("There's no fkn file")
	}

	problems := parseLines(records)

	correctAnswers := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		var userAnswer string

		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == p.a {
			correctAnswers++
		}
	}

	fmt.Printf("You got %d out of %d\n", correctAnswers, len(records))
}

type problem struct {
	q string
	a string
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
