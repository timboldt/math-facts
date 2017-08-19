package main

import (
	"flag"
	"github.com/timboldt/math-facts/challenge"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	modeString := flag.String("mode", "addition", "the type of flash card (addition, substraction, multiplication)")
	size := flag.Int("size", 9, "the biggest numbers to use, e.g. '9' in addition mode will result in problems up to 9+9")
	quantity := flag.Int("quantity", 20, "the quantity of (randomly-selected) problems to display")

	var mode challenge.Mode
	switch *modeString {
	case "addition":
		mode = challenge.AdditionMode
	case "subtraction":
		mode = challenge.SubtractionMode
	case "multiplication":
		mode = challenge.MultiplicationMode
	default:
		fmt.Printf("Invalid mode '%s'", modeString)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	trial := challenge.NewTrial(mode, *size, *quantity)
	for {
		q := trial.NextQuestion()
		if q == nil {
			break
		}
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
	}
}
