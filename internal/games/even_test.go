package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestPrepareEvenData(t *testing.T) {
	expected := [3][2]string{
		{"13", "no"},
		{"84", "yes"},
		{"20", "yes"},
	}
	actual := PrepareEvenData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
