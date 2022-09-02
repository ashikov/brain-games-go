package games

import (
	"math"
	"strconv"

	"github.com/ashikov/brain-games/src"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number == 2 || number == 3 {
		return true
	}

	for i := 2; float64(i) <= math.Sqrt(float64(number)); i ++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func PreparePrimeData() [3][2]string {
	min := 2
	max := 50

	gameData := [3][2]string{}

	for i := 0; i < roundsCount; i++ {
		number := src.GetRandomIntegerWithinRange(min, max)

		question := strconv.Itoa(number)

		var answer string
		if isPrime(number) {
			answer = "yes"
		} else {
			answer = "no"
		}

		gameData[i] = [2]string{question, answer}
	}

	return gameData
}

func RunPrime() {
	gameData := PreparePrimeData();
	desc := "Answer \"yes\" if given number is prime. Otherwise answer \"no\"."

	src.Engine(gameData, desc)
}
