package main

import "fmt"

func mileageInput() {

	totalMile := 0
	totalGallon := 0

	gasMileage := func(mile int, gallon int) float32 {
		return float32(mile) / float32(gallon)
	}

	for counter := 0; ; counter++ {

		var mileDriven, gallonUsed int

		fmt.Print("Enter mile driven (0 to cancel): ")
		fmt.Scanln(&mileDriven)
		if mileDriven == 0 {
			break
		}
		totalMile += mileDriven

		fmt.Print("Enter the gallon used: ")
		fmt.Scanln(&gallonUsed)
		totalGallon += gallonUsed
		fmt.Printf("The mileage is: %.2f\n", gasMileage(mileDriven, gallonUsed))

	}
	fmt.Printf("The total mile per gallon is: %.2f\n", gasMileage(totalMile, totalGallon))
}

func main() {
	mileageInput()
}
