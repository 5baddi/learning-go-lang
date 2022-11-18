package main

import (
	"fmt"
)

func main() {
	colors := []string{"Blue", "Purple", "Yellow"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// For loop without set the start, length and the increment
	for i := range colors {
		fmt.Println(colors[i])
	}

	// For loop key => value
	for _, color := range colors {
		fmt.Println(color)
	}

	// While
	number := 1
	for number <= 10 {
		fmt.Println("Number: ", number)
		number++
	}

	var out int
	for i := 0; i < 20; i++ {
		out = i*i + out
		if out > 12 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("The end: ", out)
}
