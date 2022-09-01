/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru

*/

package games

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1234)
}

func TestRunCalc(t *testing.T) {
	expected := [3][2]string{
		{"3 * 9", "27"},
		{"5 + 5", "10"},
		{"9 - 1", "8"},
	}
	actual := PrepareData()
	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
