package main

import "fmt"

func main() {

	var numberEntered int
	fmt.Print("Enter number: ")
	fmt.Scanln(&numberEntered)

	result := exponential(numberEntered)

	fmt.Printf("The exponetial value = %.4f", result)

}

func exponential(number int) float64 {
	exponent := 1.0

	for counter := 1; counter <= number; counter++ {
		exponent += float64(1) / float64(factorialCalc(counter))
	}
	return exponent
}
func factorialCalc(number int) int {
	var factorial = 1
	for counter := number; counter > 0; counter-- {
		factorial *= counter
	}
	return factorial
}
