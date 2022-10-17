package internal

import (
	"math/rand"
)

func GetRandomIntegerWithinRange(min, max int) int {
	return rand.Intn(max-min) + min
}
