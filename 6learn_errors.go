package main

import (
	"errors"
	"fmt"
	"os"
)

func checkInputValues(value float32) error {

	if value == 0 {
		return errors.New("value is zero")
	} else if value < 0 {
		return errors.New("value is less than zero")
	}

	return nil

}

func learnCustomErrors() {
	
	var val1, val2 float32; 
	fmt.Println("Enter the first number");
	fmt.Scan(&val1)
	
	err := checkInputValues(val1)
	if err != nil{
		panic("Val1 has the issue")
	}
	
	fmt.Println("Enter the first number");
	fmt.Scan(&val2)
	
	err1 := checkInputValues(val2)
	if err1 != nil{
		panic("Val2 has the issue")
	}

	total := val1 + val2
	result := fmt.Sprintf("%v", total)
	os.WriteFile("resultOutput.txt", [] byte(result), 0644)
}