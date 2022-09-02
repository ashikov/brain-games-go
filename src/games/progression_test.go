package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestPrepareProgressionData(t *testing.T) {
	expected := [3][2]string{
		{"38 44 50 .. 62 68 74 80 86 92", "56"},
		{"2 10 18 26 .. 42 50 58 66 74", "34"},
		{"61 .. 71 76 81 86 91 96 101 106", "66"},
	}
	actual := PrepareProgressionData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
