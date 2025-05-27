package main

import (
	"fmt"
	"sort"
)

func learnArray() {

	// arrays

	var animals [4]string = [4]string{"Lion", "Tiger", "Peacock", "Pegion"}
	fmt.Println(animals)

	// slices

	var birds = animals[2:]
	// birds = append(birds, "Parrot")  // If this statement is added then sort will modify array as well.
	fmt.Println(birds)

	sort.Slice(birds, func(i1, i2 int) bool {
		return len(birds[i1]) < len(birds[i2])
	})
	fmt.Println("Sorted slice is ", birds, animals)

	// Pointers

	var pancard *string
	pancard = new(string) // Alloactes memory and assigns it
	*pancard = "BBPIY2422K"
	fmt.Println("Pancard details", *pancard)

	// Map

	var grocery map[string]int
	grocery = make(map[string]int)
	grocery["Wheat"] = 5
	grocery["Rice"] = 5
	fmt.Println("Grocery Details", grocery)

	track := make(map[string]bool, 10)
	track["person"] = true
	fmt.Println(track)
}
