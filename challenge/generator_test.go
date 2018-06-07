package challenge_test

import (
	. "github.com/timboldt/math-facts/challenge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
)

type QuestionList []TrialQuestion

func (ql QuestionList) Len() int {
	return len(ql)
}

func (ql QuestionList) Swap(i, j int) {
	ql[i], ql[j] = ql[j], ql[i]
}

func (ql QuestionList) Less(i, j int) bool {
	return ql[i].Value1 < ql[j].Value1 ||
		(ql[i].Value1 == ql[j].Value1 &&
			ql[i].Value2 < ql[j].Value2)
}

var _ = Describe("QuestionGenerator", func() {
	Describe("constructor", func() {
		Context("a generator is created with valid inputs", func() {
			It("will create a valid generator", func() {
				g := NewGenerator(AdditionMode, 10)
				// Note: Need to account for zeroes.
				Expect(g.NumQuestions()).To(Equal(11 * 11))
				Expect(g.NumLearned()).To(Equal(0))
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
	})
	Describe("learned questions", func() {
		var (
			g *QuestionGenerator
		)
		BeforeEach(func() {
			g = NewGenerator(AdditionMode, 10)
		})
		Context("excluding a valid question that hasn't been learned yet", func() {
			It("will result in the learned question count going up", func() {
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(1))
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 2, Value2: 2, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(2))
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 0, Value2: 10, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(3))
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 10, Value2: 0, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(4))
			})
		})
		Context("excluding a valid question that has already been learned", func() {
			It("will not result in the learned question count going up", func() {
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(1))
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(1))
			})
		})
		Context("excluding a question that has a different operator", func() {
			It("will not result in the learned question count going up", func() {
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 1, Mode: SubtractionMode})
				Expect(g.NumLearned()).To(Equal(0))
			})
		})
		Context("excluding an out of range question", func() {
			It("will not result in the learned question count going up", func() {
				Expect(g.NumLearned()).To(Equal(0))
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 11, Value2: 1, Mode: AdditionMode})
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 11, Mode: AdditionMode})
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: -1, Value2: 1, Mode: AdditionMode})
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: -1, Mode: AdditionMode})
				Expect(g.NumLearned()).To(Equal(0))
			})
		})
	})
	Describe("question generation", func() {
		Context("generating a question with valid parameters", func() {
			It("will result in a question that matches those parameters", func() {
				const s = 4
				for m := range []Mode{AdditionMode, SubtractionMode, MultiplicationMode} {
					g := NewGenerator(Mode(m), s)
					n := 0
					for {
						q := g.NewQuestion()
						if q == nil {
							break
						}
						n++
						Expect(q.Mode).To(Equal(Mode(m)))
						Expect(q.Value1).To(SatisfyAll(
							BeNumerically(">=", 0),
							BeNumerically("<=", s)))
						Expect(q.Value2).To(SatisfyAll(
							BeNumerically(">=", 0),
							BeNumerically("<=", s)))
					}
					Expect(n).To(Equal(g.NumQuestions()))
				}
			})
			It("will return all unlearned questions and no learned questions", func() {
				g := NewGenerator(MultiplicationMode, 2)
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 0, Value2: 1, Mode: MultiplicationMode})
				g.ExcludeLearnedQuestion(&TrialQuestion{Value1: 1, Value2: 2, Mode: MultiplicationMode})
				var r QuestionList
				for {
					q := g.NewQuestion()
					if q == nil {
						break
					}
					r = append(r, *q)
				}
				sort.Sort(r)
				Expect(r).To(Equal(QuestionList{
					TrialQuestion{0, 0, MultiplicationMode},
					// Tuple 0,1 was excluded and should be missing.
					TrialQuestion{0, 2, MultiplicationMode},
					TrialQuestion{1, 0, MultiplicationMode},
					TrialQuestion{1, 1, MultiplicationMode},
					// Tuple 1,2 was excluded and should be missing.
					TrialQuestion{2, 0, MultiplicationMode},
					TrialQuestion{2, 1, MultiplicationMode},
					TrialQuestion{2, 2, MultiplicationMode},
				}))
			})
		})
	})
})
