package challenge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChallenge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenge Suite")
}
