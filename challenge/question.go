package challenge

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
