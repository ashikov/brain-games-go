/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru

*/

package src

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func Engine(data [3][2]string, desc string) {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Println("May I have your name? ")

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	greeting := fmt.Sprintf("Hello, %s!", strings.TrimSuffix(name, "\n"))
	fmt.Println(greeting)

	fmt.Println(desc)

	for _, v := range data {
		question := v[0]
		answer := v[1]
		questionLine := fmt.Sprintf("Question: %s", question)

		fmt.Println(questionLine)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		userAnswer := strings.TrimSuffix(input, "\n")

		if userAnswer != answer {
			placeholdersText := `'%s' is wrong answer ;(. Correct answer was '%s'."
			Let's try again, %s!`
			filledText := fmt.Sprintf(placeholdersText, userAnswer, answer, name)

			fmt.Println(filledText)
			return
		}

		fmt.Println("Correct!")
	}

	congrats := fmt.Sprintf("Congratulations, %s!", name)
	fmt.Println(congrats)
}
