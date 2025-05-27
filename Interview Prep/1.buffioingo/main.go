package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Enter the raiting?")
	element, _ := reader.ReadString('\n')
	
	element = strings.TrimSpace(element)
	raiting, err := strconv.ParseInt(element, 10, 32)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Provided raiting is ", raiting)
	}
	
}