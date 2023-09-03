package main

import "fmt"

func listRearrange(data []string) [][]string {

	var firstList []string
	var secondList []string
	var thirdList []string
	var finalList [][]string

	for counter := 0; counter < len(data); counter += 3 {
		firstList = append(firstList, data[counter])
	}

	for counter := 1; counter < len(data); counter += 3 {
		secondList = append(secondList, data[counter])
	}

	for counter := 2; counter < len(data); counter += 3 {
		thirdList = append(thirdList, data[counter])
	}

	finalList = append(finalList, firstList, secondList, thirdList)
	return finalList
}

func main() {
	var listOfLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	fmt.Println(listRearrange(listOfLetters))

}
