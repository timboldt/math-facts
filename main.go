package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/timboldt/math-facts/challenge"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	modeString = flag.String("mode", "addition", "the type of flash card (addition, substraction, multiplication)")
	mode       challenge.Mode
	size       = flag.Int("size", 9, "the biggest numbers to use, e.g. '9' in addition mode will result in problems up to 9+9")
	quantity   = flag.Int("quantity", 10, "the quantity of (randomly-selected) problems to display")
	username   = flag.String("user", "default", "username for stats purposes")
)

func init() {
	flag.Parse()

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
}

func askQuestion(q *challenge.TrialQuestion, statTracker *challenge.TrialStatTracker) {
	startTime := time.Now()
	fmt.Printf("\nHow much?  %d %s %d\n", q.Value1, q.Op, q.Value2)
	answer := getAnswer()
	if answer == q.Answer {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Oh oh! Wrong answer.")
	}
	statTracker.RecordResult(*q, challenge.TrialResult{answer == q.Answer, time.Now().Sub(startTime)})
}

func getAnswer() int {
	var answer int
	reader := bufio.NewReader(os.Stdin)
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
	return answer
}

func recordStat(question challenge.TrialQuestion, result challenge.TrialResult) {
	f, err := os.OpenFile(*username+"_stats.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := fmt.Sprintf("%s\t%d%s%d\t%v\t%.2f\n",
		time.Now().UTC().Format(time.RFC3339),
		question.Value1, question.Op, question.Value2,
		result.Correct,
		result.TimeTaken.Seconds())
	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}

func main() {
	statTracker := challenge.NewTrialStatTracker(recordStat)
	trial := challenge.NewTrial(mode, *size, *quantity)
	for {
		q := trial.NextQuestion()
		if q == nil {
			break
		}
		askQuestion(q, statTracker)
	}
	totalQuestions, totalCorrect, totalDuraction := statTracker.Summary()
	fmt.Printf("\nYou answered %d questions correctly out of %d.\nTime taken: %.1f seconds.\n", totalCorrect, totalQuestions, totalDuraction.Seconds())
}
