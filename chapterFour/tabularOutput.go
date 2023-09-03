package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%-10s%-10s%-10s%-10s\n", "N", "N2", "N3", "N4")

	for count := 1; count <= 5; count++ {
		for index := 1; index <= 4; index++ {
			fmt.Printf("%-10d", int(math.Pow(float64(count), float64(index))))
		}
		println()
	}
}
