package games

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/ashikov/brain-games/internal"
)

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

func PrepareCalcData() [3][2]string {
	min := 1
	max := 10

	operations := [3]string{"+", "-", "*"}
	gameData := [3][2]string{}

	for i := 0; i < roundsCount; i++ {
		randomIndex := rand.Intn(len(operations))
		operator := operations[randomIndex]

		a := internal.GetRandomIntegerWithinRange(min, max)
		b := internal.GetRandomIntegerWithinRange(min, max)

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
	gameData := PrepareCalcData()
	desc := "What is the result of the expression?"

	internal.Engine(gameData, desc)
}
