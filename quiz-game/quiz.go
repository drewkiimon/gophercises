package main

// go build quiz.go && ./quiz
import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const DEFAULT_FILE string = "problems.csv"
const DEFAULT_TIME_SECONDS int = 30

func main() {
	fileNameFlag := flag.String("csv", DEFAULT_FILE, "path to file you want to use")
	timeLimitFlag := flag.Int("limit", DEFAULT_TIME_SECONDS, "time limit per question")
	flag.Parse()

	f, err := os.Open(*fileNameFlag)

	if err != nil {
		log.Fatalf("We could not open the CSV file: %s", *fileNameFlag)
	}

	csvReader := csv.NewReader(f)

	lines, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Failed to parse the CSV file.")
	}

	if lines == nil {
		log.Fatal("There's no fkn file")
	}

	problems := parseLines(lines)

	correctAnswers := 0

problemLoop:
	for i, p := range problems {
		timer := time.NewTimer(time.Second * time.Duration(*timeLimitFlag))
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		// Creating a channel for a string that is asynchronous and we don't need to wait for it
		answerChannel := make(chan string)
		// go routine
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			// Send user's answer to answer channel
			answerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case userAnswer := <-answerChannel:
			if userAnswer == p.a {
				correctAnswers++
			}
		}

	}

	fmt.Printf("\nYou got %d out of %d\n", correctAnswers, len(lines))
}

// If we want to test this code, it's better to break into smaller, testable code chunks / functions
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
