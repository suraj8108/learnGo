package main

import "fmt"

func learnControlStatement() {
	var balance = 1000

	if balance > 100 {
		fmt.Println("You have a sufficient amount", balance)
	} else if balance < 50 && balance > 30 {
		fmt.Println("You are at the verge of lowest balance", balance)
	} else {
		fmt.Println("You have low balance. Please deposite amount", balance)
	}
}