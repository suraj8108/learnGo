package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// defer func() {
		// 	if r := recover(); r != nil {
		// 		fmt.Println("Recovered in goroutine:", r)
		// 	}
		// }()
		panic("Goroutine panic!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function continues...")
}
