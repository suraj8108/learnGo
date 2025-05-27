package main

import (
	"fmt"
	// "strconv"
	// "reflect"
	"sort"
)

const adhar = "1234"

func init(){
	fmt.Println("Helo")
}

func main() {

	// var name = "Suraj"
	// name, age := "Suraj", 12
	// {
	// 	age := 13
	// 	fmt.Println(age)
	// }
	// fmt.Println(name," ",age, " ", adhar)

	// fmt.Println("Enter a number")
	// fmt.Scan(&age)

	// fmt.Println("New age is ", age)
	// if age <= 10 {
	// 	fmt.Println("Younger")
	// } else if age <= 20{
	// 	fmt.Println("Smarter")
	// } else {
	// 	fmt.Println("Older")
	// }

	// toConvert := "12"
	// toage, _ := strconv.ParseInt(toConvert, 10, 64)
	// fmt.Println(toage)
	// fmt.Println(reflect.TypeOf(age))

	type Person struct {
		name string `json:"name"`
		age int `json:"age"`
	}
	
		p := new(Person)
		p.name = "Suraj"
		p.age = 12
	
		fmt.Println(p)

		// var animals []string
		animals := make([]string, 10, 10)
		animals = append(animals, "dog")
		animals = append(animals, "cat")
		animals = append(animals, "horse")
		fmt.Println(animals)

		sort.Slice(animals, func(i int, j int) bool {
			if animals[i] > animals[j] {
				return false
			} else {
				return true
			}
		})

		fmt.Println(animals)

		var ptrName *string
		fname := "Pankaj"

		ptrName = &fname

		fmt.Println(*ptrName)

		var track map[string]string
		track = map[string]string{}
		track["name"] = "Suraj";
		fmt.Println(track)
}
