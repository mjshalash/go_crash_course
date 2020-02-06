package main

// Strconv used to concatenate strings and ints
import (
	"fmt"
	"strconv"
)

//Person - Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//	Same Datatypes can also be placed on the same line
	firstName, lastName, city, gender string
	age                               int
}

// Methods do NOT go inside structs
// Two types of methods - Value and Pointer Recievers
// Value Reciever = methods that do calculations
// Pointer Recievers = methods change stuff

// Greeting method (Value Reciever, nothing is changed. Just an output)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(newLast string) {
	if p.gender == "M" {
		return
	}

	p.lastName = newLast

}

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

	// Utilizing pointer reciever method
	personOne.hasBirthday()
	personOne.getMarried("Smith")

	// Utilizing value reciever method
	fmt.Println(personOne.greet())
}
