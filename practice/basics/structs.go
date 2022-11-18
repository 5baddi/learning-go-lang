package main

import "fmt"

func main() {
	car := Car{"Renault", "Clio", 2019}

	car.ShowModel()
	car.SwitchYear(2020)

	fmt.Println(car.Year)
}

// Car is struct
type Car struct {
	Brand string
	Model string
	Year  int
}

// Show car model
func (car Car) ShowModel() {
	fmt.Println(car.Model)
}

// Swtich car year using pointer
func (car *Car) SwitchYear(year int) {
	car.Year = year
}
