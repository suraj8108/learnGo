package main

import (
	"fmt"
	"strconv"
)

func learnLoop() {
	// 4.  For Loop
	var count int
	for i := 0; i < 10; i++ {
		count += i
	}

	for count < 100 {
		fmt.Println(count)
		count += count
	}

	//5. Parsing
	intVal, _ := strconv.Atoi("10")
	fmt.Println("Parsed String to Integer", intVal)

	strVal := strconv.Itoa(10)
	fmt.Println("Parsed Integer to String", strVal)

	// 6. Memory Management make and new function
	age := new(int)
	*age = 10
	fmt.Println("Users age is ", *age)

	pincode := new(int)
	*pincode = 91
	fmt.Println("Phone pincode is +", *pincode)
}
