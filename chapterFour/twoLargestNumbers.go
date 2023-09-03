package main

import "fmt"

func main() {

	var largest = 0
	var secondLargest = 0
	var number int
	for counter := 0; counter < 10; counter++ {

		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest && number < largest {
			secondLargest = number
		}

	}

	fmt.Println("Largest number: ", largest,
		"\nSecond Largest: ", secondLargest)
}
