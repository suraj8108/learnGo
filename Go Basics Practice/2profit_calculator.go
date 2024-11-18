package main

import "fmt"

var revenue float64;
var expenses float64;
var taxRate float64;

func profitCalculator() {

	fmt.Printf("Enter total Revenue ")
	fmt.Scan(&revenue)

	fmt.Printf("Enter total Expense ")
	fmt.Scan(&expenses)

	fmt.Printf("Enter tax Rate ")
	fmt.Scan(&taxRate)

	var profitBeforeTax = calcBeforeTax(revenue, expenses)
	var profitAfterTax = calcAfterTax(profitBeforeTax, taxRate)

	ratio := profitBeforeTax / profitAfterTax

	fmt.Printf("Profit before Tax %v ", profitBeforeTax)
	fmt.Println("Profit after Tax ", profitAfterTax)

	fmt.Println("Ratio value ", ratio)
}

func calcBeforeTax(revenue, expenses float64) (float64){
	return revenue - expenses;
}
func calcAfterTax(profitBeforeTax , taxRate float64) (float64){
	return profitBeforeTax * (1 - taxRate/100);
}