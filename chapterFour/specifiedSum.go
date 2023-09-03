package main

import "fmt"

func main() {
	var number int
	var new_number int
	fmt.Print("Enter a number(pick a big one): ")
	fmt.Scanln(&number)

	var sum = 0
	for counter := 0; ; counter++ {
		fmt.Print("Enter a number to add up: ")
		fmt.Scanln(&new_number)

		sum += new_number

		if sum >= number {
			fmt.Printf("Added numbers is %d and Number you enter is %d", sum, number)
			break
		}
	}
}
