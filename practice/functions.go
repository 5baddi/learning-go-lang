package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	printHelloWorld()
	printNumber(rand.Intn(20))
	printSumTwoNumbers(rand.Intn(5), rand.Intn(10))

	sum, count := sumNumbers(rand.Intn(5), rand.Intn(10), rand.Intn(5), rand.Intn(10))
	fmt.Printf("Sum: %d for %d numbers\n", sum, count)
}

func printHelloWorld() {
	fmt.Println("Hello world!")
}

func printNumber(number int) {
	fmt.Println("Number: ", number)
}

func printSumTwoNumbers(numberOne, numberTwo int) {
	fmt.Println("Sum: ", (numberOne + numberTwo))
}

func sumNumbers(numbers ...int) (int, int) {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum, len(numbers)
}
