package main

import (
	"fmt"
)

const panNmuber = "ABCDE1234F"
var bloodGroup = "B+ve"

func main() {
	
	// Variable declaration using var
	var firstName, lastName string = "Suraj", "Yadav"

	var age, address = 25.00, "Mumbai"

	companyName := "XYZ Company"

	fmt.Println("My first name is ", firstName, " and last name is ", lastName)
	fmt.Println("My age is ", age, " and address is ", address)
	fmt.Println("My PanId is ", panNmuber, " and blood group is ", bloodGroup)
	fmt.Println("Company Name is ", companyName)

	{
		age := 10;
		fmt.Println("Age is ", age) // 10 // Create a new variable inside the block
	}
	fmt.Printf("%v age is %.0f \n", firstName, age) // 25
	msg := fmt.Sprintf("Type of age is %T", age) // 25
	fmt.Println(msg)
}