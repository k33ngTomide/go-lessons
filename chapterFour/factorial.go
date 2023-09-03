package main

import "fmt"

func main() {

	var numberEntered int
	fmt.Println("Enter number: ")
	fmt.Scanln(&numberEntered)

	result := factorial(numberEntered)
	fmt.Printf("%d! = %d", numberEntered, result)
}

func factorial(number int) int {
	var factorial = 1
	for counter := number; counter > 0; counter-- {
		factorial *= counter
	}
	return factorial
}
