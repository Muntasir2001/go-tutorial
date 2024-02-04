package main

import "fmt"

func main() {
	username := "mike"
	smsSendingLimit := 0
	costPerSms := 0.0
	hasPermission := false
	
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)
}