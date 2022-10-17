package games

import (
	"strconv"

	"github.com/ashikov/brain-games/internal"
)

func isEven(number int) bool {
	return number%2 == 0
}

func PrepareEvenData() [3][2]string {
	min := 1
	max := 100

	gameData := [3][2]string{}

	for i := 0; i < roundsCount; i++ {
		number := internal.GetRandomIntegerWithinRange(min, max)
		question := strconv.Itoa(number)

		var answer string
		if isEven(number) {
			answer = "yes"
		} else {
			answer = "no"
		}

		gameData[i] = [2]string{question, answer}
	}

	return gameData
}

func RunEven() {
	gameData := PrepareEvenData()
	desc := "Answer \"yes\" if number even otherwise answer \"no\"."

	internal.Engine(gameData, desc)
}
