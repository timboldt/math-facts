package challenge

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
	if t.asked > t.quantity {
		return nil
	}
	t.asked++
	return &TrialQuestion{Value1: 9, Value2: 9, Op: "+", Answer: 18}
}
