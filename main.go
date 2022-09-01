/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru
*/
package main

import (
	"math/rand"
	"time"

	"github.com/ashikov/brain-games/cmd"
)

func main() {
	rand.Seed(time.Now().Unix())
	cmd.Execute()
}
