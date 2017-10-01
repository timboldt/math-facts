package challenge

type Mode int

const (
	AdditionMode       = iota
	SubtractionMode
	MultiplicationMode
)

type Trial struct {
	quantity  int
	asked     int
	generator *QuestionGenerator
}

func NewTrial(mode Mode, size int, quantity int) *Trial {
	return &Trial{quantity: quantity, generator: NewGenerator(mode, size)}
}

func (t Trial) Mode() Mode {
	return t.generator.mode
}

func (t *Trial) NextQuestion() *TrialQuestion {
	if t.asked >= t.quantity {
		return nil
	}
	t.asked++
	return t.generator.NewQuestion()
}

func (t *Trial) ExcludeLearnedQuestion(q *TrialQuestion) {
	t.generator.ExcludeLearnedQuestion(q)
}

func (t *Trial) NumQuestions() int {
	return t.generator.NumQuestions()
}

func (t *Trial) NumLearnedQuestions() int {
	return t.generator.NumLearned()
}
