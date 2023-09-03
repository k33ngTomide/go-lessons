package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var binaryNumber string

	fmt.Print("Enter a binary number: ")
	fmt.Scanln(&binaryNumber)

	var newBinaryNumber, _ = strconv.Atoi(binaryNumber)
	decimalEquivalent := 0
	for counter := 0; counter < len(binaryNumber); counter++ {
		var digit = (newBinaryNumber / int(math.Pow(float64(10), float64(counter)))) % 10
		decimalEquivalent += digit * int(math.Pow(float64(2), float64(counter)))

	}
	fmt.Printf("The decimal equivalent of %s is %d\n", binaryNumber, decimalEquivalent)
}
