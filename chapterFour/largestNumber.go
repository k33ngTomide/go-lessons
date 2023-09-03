package main

import "fmt"

func main() {

	var largest = 0
	var value int
	for count := 1; count <= 10; count++ {

		fmt.Print("Enter number of sales: ")
		fmt.Scanln(&value)

		if value > largest {
			largest = value
		}

		fmt.Println("The current largest number of sale is: ", largest)
		fmt.Println("The number of count: ", count)
	}
}
