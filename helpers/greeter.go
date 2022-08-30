/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru

*/

package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func greet() {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Println("May I have your name? ")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	greeting := fmt.Sprintf("Hello, %s!", strings.TrimSuffix(text, "\n"))
	fmt.Println(greeting)
}
