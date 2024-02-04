package main

import "fmt"

func main() {
	// variables
	username := "mike"
	smsSendingLimit := 0
	costPerSms := 0.0
	hasPermission := false
	
	// %v - accepts all type
	// %d - integer
	// %f - float
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	// show type of variable
	penniesPerText := 2
	fmt.Printf("type: %T\n", penniesPerText)

	// same line var
	averateOpenRate, displayMessage := .23, "This is a message"

	fmt.Println(averateOpenRate, displayMessage)

	// conver float to int
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your account existed for", accountAgeInt, "years")

	// constant vars
	const premiumPlanName = "Premium Plan"

	// conditionals
	height := 10

	if height > 10 {
		fmt.Println("you are taller")
	} else {
		fmt.Println("you are shorter")
	}
}