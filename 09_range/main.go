package main

import "fmt"

func main() {
	// Ranges are used to loop through arrays, maps, etc.

	// Arrays/Slices
	// Create inter array of ids
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	// Underscore signals you are not using "i"
	// Unless you have underscore, error will be thrown since i is unused
	// in loop body
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmai.com"}

	// Need to account for kv pairs instead of indices
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Underscores can be used for this as well
	for _, v := range emails {
		fmt.Printf("%s\n", v)
	}
}
