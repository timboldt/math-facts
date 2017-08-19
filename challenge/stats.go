package challenge

import (
	"time"
)

type TrialResult struct {
	Correct   bool
	TimeTaken time.Duration
}

type resultList []TrialResult

type TrialStatTracker struct {
	results map[TrialQuestion]resultList
}

func NewTrialStatTracker() *TrialStatTracker {
	return &TrialStatTracker{results: make(map[TrialQuestion]resultList)}
}

func (s *TrialStatTracker) RecordResult(q TrialQuestion, r TrialResult) {
	s.results[q] = append(s.results[q], r)
}

func (s *TrialStatTracker) Summary() (quantity int, correct int, timeTaken time.Duration) {
	for _, q := range s.results {
		for _, r := range q {
			quantity++
			if r.Correct {
				correct++
			}
			timeTaken += r.TimeTaken
		}
	}
	return
}