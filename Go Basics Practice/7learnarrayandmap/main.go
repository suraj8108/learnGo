package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
)

func main() {

	var students [10]string
	students[0] = "Suraj"
	fmt.Println(students, len(students))

	var animals = []string{"Cat", "Dog"}
	fmt.Println(animals)

	//Slices
	birds := []string{"Ostrich", "Hen"}
	fmt.Println(birds)

	//arr
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		value, _ := rand.Int(rand.Reader, big.NewInt(100))
		arr[i] = int(value.Int64())
	}
	fmt.Println(arr)

	//Sort array
	sort.Strings(birds)
	fmt.Println(birds)
	// sort.Ints()
}
