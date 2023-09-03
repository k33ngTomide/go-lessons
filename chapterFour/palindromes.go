package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var number string
	var newNumber = make([]byte, 5)

	for counter := 0; ; counter++ {
		fmt.Println("Enter a five digit number: ")
		fmt.Scanln(&number)

		if len(number) != 5 {
			fmt.Println("number length is not five, try again ")
			continue
		} else {
			break
		}
	}

	for counter, letter := range number {
		newNumber[4-counter] = byte(letter)
	}
	var output = "Palindrome"
	for counter := 1; counter <= 5; counter++ {
		var anotherOne, _ = strconv.Atoi(number)
		var newAnotherOne, _ = strconv.Atoi(string(newNumber))
		var digit = anotherOne % int(math.Pow(float64(10), float64(counter))) / 10
		var newDigit = newAnotherOne % int(math.Pow(float64(10), float64(counter))) / 10

		if digit != newDigit {
			output = "Not Palindrome"
			break
		}
	}
	println(output)

}
