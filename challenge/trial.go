package challenge

import "math/rand"

type Mode int

const (
	AdditionMode       = iota
	SubtractionMode
	MultiplicationMode
)

type Trial struct {
	mode     Mode
	size     int
	quantity int
	asked    int
}

type TrialQuestion struct {
	Value1, Value2 int
	Op             string
	Answer         int
}

func NewTrial(mode Mode, size int, quantity int) *Trial {
	return &Trial{mode: mode, size: size, quantity: quantity}
}

func (t Trial) Mode() Mode {
	return t.mode
}

func (t *Trial) NextQuestion() *TrialQuestion {
	if t.asked >= t.quantity {
		return nil
	}
	t.asked++

	value1 := rand.Intn(t.size + 1)
	value2 := rand.Intn(t.size + 1)
	switch t.mode {
	case AdditionMode:
		return &TrialQuestion{Value1: value1, Value2: value2, Op: "+", Answer: value1 + value2}
	case SubtractionMode:
		// Avoid negative answers by always subtracting the smaller number from the bigger one.
		if value2 > value1 {
			value1, value2 = value2, value1
		}
		return &TrialQuestion{Value1: value1, Value2: value2, Op: "-", Answer: value1 - value2}
	case MultiplicationMode:
		return &TrialQuestion{Value1: value1, Value2: value2, Op: "*", Answer: value1 * value2}
	default:
		panic("Invalid mode.")
	}
}
