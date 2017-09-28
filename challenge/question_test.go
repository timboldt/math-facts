package challenge_test

import (
	. "github.com/timboldt/math-facts/challenge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Question", func() {
	Context("there is an addition question", func() {
		q := TrialQuestion{Value1: 3, Value2: 5, Mode: AdditionMode}
		It("will do valid arithmetic", func() {
			Expect(q.Answer()).To(Equal(8))
		})
		It("will return a valid operator", func() {
			Expect(q.Op()).To(Equal("+"))
		})
	})
	Context("there is an subtraction question", func() {
		q := TrialQuestion{Value1: 3, Value2: 5, Mode: SubtractionMode}
		It("will do valid arithmetic", func() {
			Expect(q.Answer()).To(Equal(-2))
		})
		It("will return a valid operator", func() {
			Expect(q.Op()).To(Equal("-"))
		})
	})
	Context("there is an multiplication question", func() {
		q := TrialQuestion{Value1: 3, Value2: 5, Mode: MultiplicationMode}
		It("will do valid arithmetic", func() {
			Expect(q.Answer()).To(Equal(15))
		})
		It("will return a valid operator", func() {
			Expect(q.Op()).To(Equal("*"))
		})
	})
	Context("there is an invalid question", func() {
		q := TrialQuestion{Value1: 3, Value2: 5, Mode: Mode(99)}
		It("will panic", func() {
			Expect(func() { _ = q.Answer() }).Should(Panic())
		})
		It("will panic", func() {
			Expect(func() { _ = q.Op() }).Should(Panic())
		})
	})
})
