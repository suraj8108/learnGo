package main

import (
	"fmt"
	"os"
)

func makeTheMessageEntry(msg string){
	os.WriteFile("messageSwitch.txt", []byte(msg), 0644)
}

func learnLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Current Iterations is", i)
	}

	fmt.Println("Loop Ended Successfully")
	for {
		var choice int;
		fmt.Println("Please Enter the choice");
		fmt.Println("1. Good Morning");
		fmt.Println("2. Good Afternoon");
		fmt.Println("3. Good Evening");

		fmt.Scan(&choice)
		var msg string
		switch choice{
			case 1:
				msg = "Morning 6AM to Afternoon 12.00 AM"
				fmt.Print(msg)
				makeTheMessageEntry(msg);
			case 2:
				msg = "Afternoon 12.00 PM to Evening 4.00 PM"
				fmt.Print(msg)
				makeTheMessageEntry(msg);
			case 3:
				msg = "Evening 4.00 PM to Night until you sleep"
				fmt.Print(msg)
				makeTheMessageEntry(msg);
			default: 
			fmt.Println("Good Bye")
			return
		}

		
	}
}