package games

import (
	"fmt"
	"strconv"

	"github.com/ashikov/brain-games/internal"
)

func getGcd(a, b int) int {
	if b != 0 {
		return getGcd(b, a%b)
	}

	return a
}

func PrepareGcdData() [3][2]string {
	min := 2
	max := 50

	gameData := [3][2]string{}

	for i := 0; i < roundsCount; i++ {
		number1 := internal.GetRandomIntegerWithinRange(min, max)
		number2 := internal.GetRandomIntegerWithinRange(min, max)

		question := fmt.Sprintf(
			"%s %s",
			strconv.Itoa(number1),
			strconv.Itoa(number2),
		)

		gcd := getGcd(number1, number2)

		answer := strconv.Itoa(gcd)

		gameData[i] = [2]string{question, answer}
	}

	return gameData
}

func RunGcd() {
	gameData := PrepareGcdData()
	desc := "Find the greatest common divisor of given numbers."

	internal.Engine(gameData, desc)
}
