package challenge_test

import (
	. "github.com/timboldt/math-facts/challenge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuestionGenerator", func() {
	var (
		g *QuestionGenerator
	)
	BeforeEach(func() {
		g = NewGenerator(AdditionMode, 10)
	})
	Context("a generator is created with valid inputs", func() {
		It("will create a valid generator", func() {
			// Note: Need to account for zeroes.
			Expect(g.NumQuestions()).To(Equal(11 * 11))
			Expect(g.NumBanned()).To(Equal(0))
		})
	})
	Context("a generator is created with an invalid mode", func() {
		It("will panic", func() {
			Expect(func() { _ = NewGenerator(Mode(99), 10) }).Should(Panic())
		})
	})
	Context("a generator is created with an invalid size", func() {
		It("will panic", func() {
			Expect(func() { _ = NewGenerator(AdditionMode, 999) }).Should(Panic())
			Expect(func() { _ = NewGenerator(AdditionMode, 0) }).Should(Panic())
			Expect(func() { _ = NewGenerator(AdditionMode, -1) }).Should(Panic())
		})
	})
	Context("banning a valid question that hasn't been banned yet", func() {
		It("will result in the ban count going up", func() {
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
			Expect(g.NumBanned()).To(Equal(1))
			g.BanQuestion(&TrialQuestion{Value1: 2, Value2: 2, Mode: AdditionMode})
			Expect(g.NumBanned()).To(Equal(2))
		})
	})
	Context("banning a valid question that has already been banned", func() {
		It("will not result in the ban count going up", func() {
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
			Expect(g.NumBanned()).To(Equal(1))
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
			Expect(g.NumBanned()).To(Equal(1))
		})
	})
	Context("banning a question that has a different operator", func() {
		It("will not result in the ban count going up", func() {
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: SubtractionMode})
			Expect(g.NumBanned()).To(Equal(0))
		})
	})
	Context("banning a out of range question", func() {
		It("will not result in the ban count going up", func() {
			g := NewGenerator(AdditionMode, 10)
			Expect(g.NumBanned()).To(Equal(0))
			g.BanQuestion(&TrialQuestion{Value1: 11, Value2: 1, Mode: AdditionMode})
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: 11, Mode: AdditionMode})
			g.BanQuestion(&TrialQuestion{Value1: -1, Value2: 1, Mode: AdditionMode})
			g.BanQuestion(&TrialQuestion{Value1: 1, Value2: -1, Mode: AdditionMode})
			Expect(g.NumBanned()).To(Equal(0))
		})
	})
})

/***

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

func (g *QuestionGenerator) NumBanned() int {
	return len(g.banned)
}

 */
