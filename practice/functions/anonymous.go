package main

import "fmt"

func nextValue() func() int {
	number := 0

	return func() int {
		number++

		return number
	}
}

func main() {
	nextNumber := nextValue()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	newNumber := nextValue()
	fmt.Println(newNumber())
}
