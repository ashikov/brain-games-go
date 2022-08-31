/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru

*/

package src

import (
	"math/rand"
	"time"
)

func GetRandomIntegerWithinRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}
