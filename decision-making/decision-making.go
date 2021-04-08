package main

import (
	"fmt"
	"math/rand"
	"time"
)

func returnsTwoThings() (string, bool) {
	return "Nothing", rand.Intn(2) == 1
}

func getRandomNumber(a int, b int) int {
	return rand.Intn(b-a) + a
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// This randomly returns true
	if _, ok := returnsTwoThings(); ok {
		fmt.Println("Something")
	}

	randomNumber := getRandomNumber(1, 4)
	switch randomNumber {
	case 1:
		fmt.Println("Numero Uno")
	case 2:
		fmt.Println("It is 222")
	default:
		fmt.Println("What is it???")
	}
}
