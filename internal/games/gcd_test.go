package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestPrepareGcdData(t *testing.T) {
	expected := [3][2]string{
		{"14 2", "2"},
		{"45 30", "15"},
		{"35 45", "5"},
	}
	actual := PrepareGcdData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
