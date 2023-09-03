package main

import "fmt"

func main() {

	var accountNumber, beginningBalance, itemCharges, creditApplied, creditLimit int

	for counter := 0; ; counter++ {

		accountNumber = inputCollect("Enter account number of customer (enter 0 to cancel):  ")
		if accountNumber == 0 {
			break
		}

		beginningBalance = inputCollect("Enter the balance at the beginning of the Month :  ")
		itemCharges = inputCollect("Enter the item charges for the Month: ")
		creditApplied = inputCollect("Enter the credit applied to the customer: ")
		creditLimit = inputCollect("Enter customer credit limit: ")

		newBalance := beginningBalance + itemCharges - creditApplied

		if newBalance > creditLimit {
			fmt.Println("Credit Limit Exceeded")
		}
	}

}

func inputCollect(message string) int {
	var variable int
	fmt.Print(message)
	fmt.Scanln(&variable)

	return variable
}
