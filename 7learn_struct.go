package main

import (
	"fmt"

	"github.com/suraj8108/learnGo/user"
)



func learnStruct() { 
	fmt.Println("Enter firstName")
	firstName := getUserInput()

	fmt.Println("Enter lastName")
	lastName := getUserInput()

	fmt.Println("Enter Birth Date")
	birthDate := getUserInput()

	userDetails := user.New(firstName, lastName, birthDate)

	adminDetails := user.NewAdmin("suraj@admin.com", "1234", userDetails)
	
	fmt.Println(adminDetails)
	
	adminDetails.Student.PrintUserDetails()
	// adminDetails.ClearUserNames()
	adminDetails.User.PrintUserDetails()

}

func getUserInput() (input string) {
	fmt.Scan(&input)
	return input
}

