package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var number int = rand.Intn(5)
	var message string

	if number > 3 {
		message = "Wow"
	} else if number > 1 {
		message = "Pff"
	} else {
		message = "Oops"
	}

	fmt.Printf("%s you get %dk $\n", message, number)
}
