package challenge

import (
	"math/rand"
	"fmt"
	"os"
)

type TrialQuestion struct {
	Value1, Value2 int
	Mode           Mode
}

func (q TrialQuestion) Answer() int {
	switch q.Mode {
	case AdditionMode:
		return q.Value1 + q.Value2
	case SubtractionMode:
		return q.Value1 - q.Value2
	case MultiplicationMode:
		return q.Value1 * q.Value2
	default:
		panic("Invalid mode.")
	}
}

func (q TrialQuestion) Op() string {
	switch q.Mode {
	case AdditionMode:
		return "+"
	case SubtractionMode:
		return "-"
	case MultiplicationMode:
		return "*"
	default:
		panic("Invalid mode.")
	}
}

type QuestionGenerator struct {
	mode   Mode
	size   int
	seq    []int
	next   int
	banned map[TrialQuestion]bool
}

func NewGenerator(mode Mode, size int) *QuestionGenerator {
	if size > 30 {
		fmt.Println("Maximum size is 30.")
		os.Exit(1)
	}
	return &QuestionGenerator{mode: mode, size: size, seq: rand.Perm(size * size), banned: make(map[TrialQuestion]bool)}
}

func (g *QuestionGenerator) NewQuestion() *TrialQuestion {
	for {
		if g.next >= len(g.seq) {
			return nil
		}

		value1 := g.seq[g.next] / g.size
		value2 := g.seq[g.next] % g.size
		// Avoid negative answers by always subtracting the smaller number from the bigger one.
		if g.mode == SubtractionMode && value2 > value1 {
			value1, value2 = value2, value1
		}
		g.next++

		q := &TrialQuestion{Value1: value1, Value2: value2, Mode: g.mode}
		if !g.banned[*q] {
			return q
		}
	}
}

func (g *QuestionGenerator) BanQuestion(q *TrialQuestion) {
	g.banned[*q] = true
}
