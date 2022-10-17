package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestPrepareCalcData(t *testing.T) {
	expected := [3][2]string{
		{"3 * 9", "27"},
		{"5 + 5", "10"},
		{"9 - 1", "8"},
	}
	actual := PrepareCalcData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
