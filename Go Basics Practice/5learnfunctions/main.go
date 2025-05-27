package main

import "fmt"

func add(a int, b int) (string, int) {
	sum := a + b
	result := fmt.Sprintf("Sum of %v and %v is ", a, b)
	return result, sum
}

var subtract = func(a int, b int) int {
	return a - b
}

var multiply = func(values ...int) (result int) {
	result = 1
	for _, element := range values {
		result *= element
	}
	return result
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(subtract(10, 20))
	fmt.Println(multiply(10, 20, 30))

}
