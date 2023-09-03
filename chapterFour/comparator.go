package main

import "fmt"

func main() {
	var number int
	var new_number int

	for i := 0; ; i++ {

		fmt.Print("Enter a number (enter 0 to cancel): ")
		fmt.Scanln(&number)
		if number == 0 {
			break
		}

		fmt.Print("Enter a number: ")
		fmt.Scanln(&new_number)

		if new_number > number {
			print("-1")
		} else if number > new_number {
			print("1")
		} else if number == new_number {
			print("0")
		}

	}

}
