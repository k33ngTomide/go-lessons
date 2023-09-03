package main

import (
	"fmt"
	"math"
)

func main() {

	var numberEntered int
	fmt.Print("Enter number: ")
	fmt.Scanln(&numberEntered)

	result := exponentialPower(numberEntered)

	fmt.Printf("The exponetial value = %.4f", result)

}

func exponentialPower(number int) float64 {
	exponent := 1.0

	for counter := 1; counter <= number; counter++ {
		exponent += math.Pow(float64(number), float64(counter)) / float64(factorialCalculator(counter))
	}
	return exponent
}
func factorialCalculator(number int) int {
	var factorial = 1
	for counter := number; counter > 0; counter-- {
		factorial *= counter
	}
	return factorial
}
