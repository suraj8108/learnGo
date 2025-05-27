package main

import "fmt"

type Person struct {
	Name    string `json:"name"`
	College string `json:"college"`
	Age     int    `json:"age"`
}

func (p *Person) updateName(name string) {
	p.Name = name
}

func NewPerson(name string, college string, age int) *Person {
	return &Person{
		name,
		college,
		age,
	}
}

func main() {
	p1 := Person{"Suraj", "Thakur", 25}
	p2 := Person{
		Age:     21,
		Name:    "Ram",
		College: "Nirmala",
	}
	fmt.Println(p1, p2)

	p3 := NewPerson("Pankaj", "MBA", 26)
	fmt.Println(p3)
	p3.updateName("Pankaj Patil")
	fmt.Println(p3)
}
