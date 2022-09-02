package games

import (
	"strconv"
	"strings"

	"github.com/ashikov/brain-games/src"
)

func PrepareProgressionData() [3][2]string {
	gameData := [3][2]string{}
	length := 10

	for i := 0; i < roundsCount; i++ {
		start := src.GetRandomIntegerWithinRange(1, 70)
		step := src.GetRandomIntegerWithinRange(2, 10)
		hiddenElementIndex := src.GetRandomIntegerWithinRange(0, length - 1)

		progression := [10]string{}

		for j := 0; j < length; j++ {
			if j == hiddenElementIndex {
				progression[j] = ".."
			} else {
				progression[j] = strconv.Itoa(start + step * j)
			}
		}

		question := strings.Join(progression[:], " ")
		answer := strconv.Itoa(start + step * hiddenElementIndex)

		gameData[i] = [2]string{question, answer}
	}

	return gameData
}

func RunProgression() {
	gameData := PrepareProgressionData()
	desc := "What number is missing in the progression?"

	src.Engine(gameData, desc)
}
