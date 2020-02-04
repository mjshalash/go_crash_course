package main

import "fmt"

func main() {
	// Pointers are used to point to memory address of a variable
	a := 5
	b := &a // b is pointing to memory address of variable a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // --> Returns *int

	// Using * to read value stored at address
	fmt.Println(*b)  // --> Returns 5
	fmt.Println(*&a) // --> also returns 5

	// Can change value using address pointer
	*b = 10
	fmt.Println(a) //--> Outputs 10

	// Everything is GO is passed by value
}
