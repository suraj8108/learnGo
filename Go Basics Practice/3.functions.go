package main

import "fmt"

const middleName = "R."
func learnFunc() {
	fullName, issuePresent := printDetails("Suraj", "Yadav")

	if !issuePresent {
		fmt.Print(fullName)
	}
}

func printDetails(fname, lname string) (string, bool){
	anyIssue := false
	return fmt.Sprintf("%v %v %v", fname, middleName, lname) , anyIssue
}

func printDetailsNoReturn(fname, lname string) (fullName string, anyIssue bool){
	anyIssue = false
	fullName = fmt.Sprintf("%v %v %v", fname, middleName, lname)
	return
}