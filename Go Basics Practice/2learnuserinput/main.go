package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using Scan function

	var name string

	fmt.Println("Enter your name")
	fmt.Scan(&name) // Takes the single string as input is space present the content before space is only stored.

	// Use bufio
	file, _ := os.Open("data.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Will fetch the college name from the file.")
	college, _ := reader.ReadString('\n')

	fmt.Printf("%v has went to %v college", name, college)

}
