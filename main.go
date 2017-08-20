package main

import (
	"flag"
	"github.com/timboldt/math-facts/challenge"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

func main() {
	modeString := flag.String("mode", "addition", "the type of flash card (addition, substraction, multiplication)")
	size := flag.Int("size", 9, "the biggest numbers to use, e.g. '9' in addition mode will result in problems up to 9+9")
	quantity := flag.Int("quantity", 10, "the quantity of (randomly-selected) problems to display")
	flag.Parse()

	var mode challenge.Mode
	switch *modeString {
	case "addition":
	case "add":
		mode = challenge.AdditionMode
	case "subtraction":
	case "sub":
		mode = challenge.SubtractionMode
	case "multiplication":
	case "mult":
		mode = challenge.MultiplicationMode
	default:
		fmt.Printf("Invalid mode '%s'", modeString)
		os.Exit(1)
	}
	rand.Seed(time.Now().UTC().UnixNano())

	reader := bufio.NewReader(os.Stdin)
	statTracker := challenge.NewTrialStatTracker()
	trial := challenge.NewTrial(mode, *size, *quantity)
	for {
		q := trial.NextQuestion()
		if q == nil {
			break
		}
		startTime := time.Now()
		fmt.Printf("\nHow much?  %d %s %d\n", q.Value1, q.Op, q.Value2)
		var answer int
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				//fmt.Println("%v %v", []byte(line), err)
				fmt.Println("Please enter a number.")
				continue
			}
			answer, err = strconv.Atoi(strings.Trim(line, "\r\n"))
			if err != nil {
				//fmt.Println(err)
				fmt.Println("Please enter a number.")
				continue
			}
			break
		}
		if answer == q.Answer {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Oh oh! Wrong answer.")
		}
		statTracker.RecordResult(*q, challenge.TrialResult{answer == q.Answer, time.Now().Sub(startTime)})
	}
	totalQuestions, totalCorrect, totalDuraction := statTracker.Summary()
	fmt.Printf("\nYou answered %d questions correctly out of %d.\nTime taken: %.1f seconds.\n", totalCorrect, totalQuestions, totalDuraction.Seconds())
}
