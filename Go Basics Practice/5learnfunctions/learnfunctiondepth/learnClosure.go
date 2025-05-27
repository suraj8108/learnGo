package learnfunctiondepth

import "fmt"

func LearnClosure() {
	elements := [4]int{1, 2, 3, 4}
	double := translate(2)
	fmt.Println(tranform(elements, double))

	triple := translate(3)
	fmt.Println(tranform(elements, triple))
}

func translate(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func tranform(elements [4]int, translateLogic func(int) int) []int {
	result := make([]int, 4)
	for index, element := range elements {
		result[index] = translateLogic(element)
	}

	return result
}
