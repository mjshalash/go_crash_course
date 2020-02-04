package main

import "fmt"

//Person - Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Methods do NOT go inside structs
// Two types of methods - Value and Pointer Recievers
// Value Reciever = methods that do calculations
// Pointer Recievers = methods change stuff

func main() {
	// Structs are one of the most important things in GO
	// Similar to classes
	// personOne := Person{firstName: "Malik", lastName: "Shalash", city: "Boulder", gender: "M", age: 25}

	// Do not have to use field labels
	personOne := Person{"Malik", "Shalash", "Boulder", "M", 25}

	// fmt.Println(personOne)

	// Get single value
	fmt.Println(personOne.age)

	// Modify value
	personOne.age++
	fmt.Println(personOne.age)
}
