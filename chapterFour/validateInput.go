package main

import "fmt"

func main() {

	var input int
	for true {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&input)

		if input == 1 || input == 2 {
			fmt.Println("The number you entered: ", input)
			break
		}

	}
}
