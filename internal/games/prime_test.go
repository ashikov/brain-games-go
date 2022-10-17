package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestPreparePrimeData(t *testing.T) {
	expected := [3][2]string{
		{"21", "no"},
		{"35", "no"},
		{"26", "no"},
	}
	actual := PreparePrimeData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
