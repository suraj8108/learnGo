package main

import (
	"fmt"
	"time"
)

func sayHi(){
	fmt.Println("Hi Suraj")
	for i  := 0; i < 10; i++{
		fmt.Println("Hii", i)
		time.Sleep(1000*time.Millisecond)
	}
	fmt.Println("Hi Suraj Ended")

}

func sayHello(){
	fmt.Println("Hello Suraj")
	for i := 0; i < 10; i++{
		fmt.Println("Hello", i)
		time.Sleep(1000*time.Millisecond)
	}
	fmt.Println("Hello Suraj Ended")
}

func main(){
	fmt.Println("Learn GO Routines");

	// go sayHi();
	// go sayHello();
	// time.Sleep(4000*time.Millisecond)

	
}