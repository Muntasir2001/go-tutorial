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
	
	// conditional smaller version
	/*
		if INITIAL_STATEMENT; CONDITION {

		}
	*/
	if length := 12; length < 1 {
		fmt.Println("Number is invalid")
	}

	// functions
	fmt.Println(concat("Hello", "World"))

	// variables in Go variables are passed by value, not by reference
	// eg:
	sendsSoFar := 430
	const sendsToAdd = 25

	// to update sendsSoFar, sendsSoFar needs to be updated/reassigned to be updated
	sendsSoFar = incrementSends(430, 25)
	fmt.Println("you've sent", sendsSoFar, "messages")

	// ignore a return value/variable (since golang doesn't allow unused variables)
	firstName, _ := getNames()
	fmt.Println(firstName)

	// named return values
	x, y := getCords()
	fmt.Println(x, y)
}

// functions
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

// multiple return value
func getNames() (string, string) {
	return "John", "Doe"
}

// named return values
func getCords() (x int, y int) {
	x = 12
	y = 1222
	
	return
}