package main

import "fmt"

func main() {
	taxCalculator()
}

func taxCalculator() {
	var value float64
	var amount float64
	var taxRate float64

	for count := 0; ; count++ {

		fmt.Print("Enter earnings (Enter -1 to stop): ")
		fmt.Scanln(&value)

		if value == -1 {
			break
		}

		if value > 30000 {
			taxRate = 0.20

		} else if value <= 30000 {
			taxRate = 0.15
		}
		amount = value * taxRate
		fmt.Printf("The tax payment is: %.2f\n", amount)

	}
}
