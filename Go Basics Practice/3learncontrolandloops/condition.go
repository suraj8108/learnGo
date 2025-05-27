package main

import (
	"fmt"
)

func learnCondition() {

	// Conditions in GoLang
	var marks int = 10
	if marks < 50 {
		fmt.Println("Good")
	} else if marks < 70 {
		fmt.Println("Excellent")
	} else {
		fmt.Println("Master Blasters")
	}

	for i := 10; i < 10; i++ {
		fmt.Println(i)
	}

	// Swith Statement
	//1. Using OptionalStatement

	switch day := "SUNDAY"; day {
	case "SUNDAY":
		fmt.Println("Its a holiday")
	case "SATURDAY":
		fmt.Println("Its a holiday")
	default:
		fmt.Println("Its a Working Day")
	}

	//2. Using OptionalExpression
	switch {
	case marks < 60:
		fmt.Println("Good")
	case marks < 70:
		fmt.Println("Excellent")
	default:
		fmt.Println("Master Blasters")
	}

	// 3. Using Type Switch
	var amount interface{} = 1000.12
	switch amount.(type) {
	case int:
		fmt.Println("It is an integer")
	case float32:
		fmt.Println("It is an float 32")
	case float64:
		fmt.Println("It is an float 64")
	default:
		fmt.Println("Unknown Type")
	}

}
