package challenge

import (
	"math/rand"
)

type QuestionGenerator struct {
	mode    Mode
	size    int
	seq     []int
	next    int
	learned map[TrialQuestion]bool
}

func NewGenerator(mode Mode, size int) *QuestionGenerator {
	switch mode {
	case AdditionMode:
	case SubtractionMode:
	case MultiplicationMode:
		break
	default:
		panic("Invalid mode.")
	}
	if size > 30 {
		panic("Maximum size is 30.")
	}
	if size < 1 {
		panic("Minimum size is 1.")
	}
	// Need to account for zeroes.
	totalSize := (size + 1) * (size + 1)
	return &QuestionGenerator{mode: mode, size: size, seq: rand.Perm(totalSize), learned: make(map[TrialQuestion]bool)}
}

func (g *QuestionGenerator) NewQuestion() *TrialQuestion {
	for {
		if g.next >= len(g.seq) {
			return nil
		}

		value1 := g.seq[g.next] / (g.size + 1)
		value2 := g.seq[g.next] % (g.size + 1)
		g.next++

		q := &TrialQuestion{Value1: value1, Value2: value2, Mode: g.mode}
		if !g.learned[*q] {
			return q
		}
	}
}

func (g *QuestionGenerator) NumQuestions() int {
	return len(g.seq)
}

func (g *QuestionGenerator) ExcludeLearnedQuestion(q *TrialQuestion) {
	if q.Value1 >= 0 && q.Value1 <= g.size &&
		q.Value2 >= 0 && q.Value2 <= g.size &&
		q.Mode == g.mode {
		g.learned[*q] = true
	}
}

func (g *QuestionGenerator) NumLearned() int {
	return len(g.learned)
}
