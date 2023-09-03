package main

import "fmt"

func main() {
	for counter := 0; counter <= 10; counter++ {
		for index := 0; index <= counter; index++ {
			fmt.Print("*")
		}
		for index := 0; index < 11-counter; index++ {
			fmt.Print(" ")
		}
		for index := 0; index < 11-counter; index++ {
			fmt.Print("*")
		}
		for index := 0; index <= counter; index++ {
			fmt.Print(" ")
		}
		for index := 0; index <= counter; index++ {
			fmt.Print(" ")
		}
		for index := 0; index < 11-counter; index++ {
			fmt.Print("*")
		}
		for index := 0; index < 11-counter; index++ {
			fmt.Print(" ")
		}
		for index := 0; index <= counter; index++ {
			fmt.Print("*")
		}
		println()
	}
}
