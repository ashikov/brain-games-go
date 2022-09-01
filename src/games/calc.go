package games

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/ashikov/brain-games/src"
)


const desc = "What is the result of the expression?"
const RoundsCount = 3
const min = 1
const max = 10

func calculate(operator string, number1, number2 int) (int, error) {
	switch operator {
		case "+":
			return number1 + number2, nil
		case "-":
			return number1 - number2, nil
		case "*":
			return number1 * number2, nil
		default:
			error := fmt.Sprintf("Unknown operator: %s", operator)
			return 0, errors.New(error)
	}
}

func PrepareData() [3][2]string {
	operations := [3]string{"+", "-", "*"}
	gameData := [3][2]string{}

	for i := 0; i < RoundsCount; i++ {
		randomIndex := rand.Intn(len(operations))
		operator := operations[randomIndex]

		a := src.GetRandomIntegerWithinRange(min, max)
		b := src.GetRandomIntegerWithinRange(min, max)

		question := fmt.Sprintf(
			"%s %s %s",
			strconv.Itoa(a),
			operator,
			strconv.Itoa(b),
		)

		result, _ := calculate(operator, a, b)
		correctAnswer := strconv.Itoa(result)

		gameData[i] = [2]string{question, correctAnswer}
	}

	return gameData
}

func RunCalc() {
	gameData := PrepareData()
	src.Engine(gameData, desc)
}
