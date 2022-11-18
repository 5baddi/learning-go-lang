package main

import "fmt"

func fact(number int) int {
	if number == 0 {
		return 1
	}

	return number * fact(number-1)
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("The factorial of %d is: %d\n", number, fact(number))
}
