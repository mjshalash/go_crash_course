package main

import "fmt"

func main() {
	// Define map (key-value pairs)
	// Format - map[key dtype]value dtype
	// emails := make(map[string]string)

	// Assign Key Value Pairs
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Can declare and assign at same time
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmai.com"}

	// Output Items
	//fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete Items
	delete(emails, "Bob")
	fmt.Println(emails)
}
