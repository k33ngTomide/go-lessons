package main

import (
	"fmt"
	"strconv"
)

func main() {

	var amount = calculatePayment()
	fmt.Println("Week payment: ", amount)
}

func calculatePayment() float64 {
	var itemValue float64
	var amount float64
	var newItem string
	total := 0.0
	for count := 0; ; count++ {
		itemValue, _ = strconv.ParseFloat(inputCollector("Enter item value: "), 64)

		total += itemValue
		newItem = inputCollector("Add more items?  Enter \n1 to Add more  \n2 to stop  \nEnter: ")

		if newItem == "1" {
			continue
		} else {
			amount = (0.09 * float64(total)) + 200
			break
		}
	}
	return amount
}

func inputCollector(message string) string {
	var variable string
	fmt.Print(message)
	fmt.Scanln(&variable)

	return variable
}
