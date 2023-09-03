package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var userPin string
	fmt.Print("Enter your pin: ")
	fmt.Scanln(&userPin)

	newEncryptedPin := encrytPin(userPin)
	fmt.Printf("The encrypted pin is: %s\n", newEncryptedPin)

	formerUserPin := decrytPin(newEncryptedPin)
	fmt.Printf("The decrypted of the enceypted pin is: %s\n", formerUserPin)

}

func encrytPin(pin string) string {
	const seven = 7
	const ten = 10
	newPin, _ := strconv.Atoi(pin)
	var newDigitStr = ""
	for counter := len(pin) - 1; counter >= 0; counter-- {
		digit := (newPin / int(math.Pow(float64(ten), float64(counter)))) % ten

		newDigit := (digit + seven) % ten
		digitStr := strconv.Itoa(newDigit)

		newDigitStr = newDigitStr + digitStr

	}
	var encryptedPassword string
	encryptedPassword += string(newDigitStr[2])
	encryptedPassword += string(newDigitStr[3])
	encryptedPassword += string(newDigitStr[0])
	encryptedPassword += string(newDigitStr[1])

	return encryptedPassword
}
func decrytPin(pin string) string {

	var decryptedPassword string
	decryptedPassword += string(pin[2])
	decryptedPassword += string(pin[3])
	decryptedPassword += string(pin[0])
	decryptedPassword += string(pin[1])

	const seven = 7
	const ten = 10
	newPin, _ := strconv.Atoi(decryptedPassword)
	var newDigitStr = ""

	for counter := len(pin) - 1; counter >= 0; counter-- {
		digit := (newPin / int(math.Pow(float64(ten), float64(counter)))) % ten

		var newDigit int
		if digit >= seven {
			newDigit = digit - seven
		} else {
			newDigit = digit - seven + ten
		}
		digitStr := strconv.Itoa(newDigit)

		newDigitStr += digitStr

	}
	return newDigitStr
}
