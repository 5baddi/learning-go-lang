package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var message string
	var day int = rand.Intn(7) + 1

	switch day {
	case 1:
		message = "It's Sunday!"
	case 2:
		message = "It's Monday!"
	default:
		message = "It's some other day!"
	}

	fmt.Println(message, day)
}
