/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru

*/

package src

import (
	"math/rand"
)

func GetRandomIntegerWithinRange(min, max int) int {
	return rand.Intn(max - min) + min
}
