package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/timboldt/math-facts/challenge"
)

var (
	modeString = flag.String("mode", "addition", "the type of flash card (addition, substraction, multiplication)")
	mode       challenge.Mode
	sizeFlag   = flag.Int("size", 9, "the biggest numbers to use, e.g. '9' in addition mode will result in problems up to 9+9")
	quantity   = flag.Int("quantity", 10, "the quantity of (randomly-selected) problems to display")
	username   = flag.String("user", "default", "username for stats purposes")
)

func init() {
	flag.Parse()

	switch *modeString {
	case "addition", "add":
		mode = challenge.AdditionMode
	case "subtraction", "sub":
		mode = challenge.SubtractionMode
	case "multiplication", "mult":
		mode = challenge.MultiplicationMode
	default:
		fmt.Printf("Invalid mode %q\n", *modeString)
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())
}

func askQuestion(q *challenge.TrialQuestion, statTracker *challenge.TrialStatTracker) {
	startTime := time.Now()
	answer := getAnswer(q)
	if answer == q.Answer() {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Oh oh! Wrong answer.")
	}
	statTracker.RecordResult(*q, challenge.TrialResult{
		Correct:   answer == q.Answer(),
		TimeTaken: time.Since(startTime),
	})
}

func getAnswer(q *challenge.TrialQuestion) int {
	var answer int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	for {
		fmt.Printf("How much?  %d %s %d\n", q.Value1, q.Op(), q.Value2)
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
		question.Value1, question.Op(), question.Value2,
		result.Correct,
		result.TimeTaken.Seconds())
	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}

func processStats(trial *challenge.Trial) {
	f, err := os.Open(*username + "_stats.txt")
	if err != nil {
		fmt.Println("Unable to read existing stats file. Ignoring error.")
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	re := regexp.MustCompile(`\t(\d+)([\+\-\*])(\d+)\t(true|false)\t(\d+)`)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		m := re.FindStringSubmatch(s)
		if len(m) < 6 {
			if len(s) > 5 {
				fmt.Printf("Invalid stat file string: %s\n", s)
			}
			continue
		}
		val1, err := strconv.Atoi(m[1])
		if err != nil {
			fmt.Printf("Invalid operand %s\n", m[1])
			continue
		}
		val2, err := strconv.Atoi(m[3])
		if err != nil {
			fmt.Printf("Invalid operand %s\n", m[3])
			continue
		}
		var mode challenge.Mode
		switch m[2] {
		case "+":
			mode = challenge.AdditionMode
		case "-":
			mode = challenge.SubtractionMode
		case "*":
			mode = challenge.MultiplicationMode
		default:
			fmt.Printf("Invalid operator %s\n", m[2])
			continue
		}
		correct, err := strconv.ParseBool(m[4])
		if err != nil {
			fmt.Printf("Invalid bool %s\n", m[4])
			continue
		}
		seconds, err := strconv.Atoi(m[5])
		if err != nil {
			fmt.Printf("Invalid time %s\n", m[5])
			continue
		}

		if correct && seconds < 8 && mode == trial.Mode() {
			trial.ExcludeLearnedQuestion(&challenge.TrialQuestion{Value1: val1, Value2: val2, Mode: mode})
		}
	}
}

func main() {
	statTracker := challenge.NewTrialStatTracker(recordStat)
	trial := challenge.NewTrial(mode, *sizeFlag, *quantity)
	// Exclude any questions that were too easy.
	processStats(trial)
	// Exclude negative answers.
	skipped := 0
	if trial.Mode() == challenge.SubtractionMode {
		for i := 0; i <= *sizeFlag; i++ {
			for j := i + 1; j <= *sizeFlag; j++ {
				// TODO: It would be useful if the generator distinguished between questions that were learned vs ones
				// that were excluded for other reasons (e.g. negatives). This would eliminate all the extra "skipped" logic.
				trial.ExcludeLearnedQuestion(&challenge.TrialQuestion{Value1: i, Value2: j, Mode: mode})
				skipped++
			}
		}
	}
	learned := trial.NumLearnedQuestions() - skipped
	total := trial.NumQuestions() - skipped
	fmt.Printf("Known questions: %d (%d%%)\n", learned, 100*learned/total)
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
