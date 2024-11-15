package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) *User {
	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}
}

func (u *User) PrintUserDetails() {
	fmt.Printf("User details %v", u)
}

func (u *User) ClearUserNames() {
	u.firstName = ""
	u.lastName = ""
}

type Student struct {
	studentName string
}

func (s *Student) PrintUserDetails(){
	fmt.Printf("Student details %v", s)
}

type Admin struct {
	email string
	pass string
	User 
	Student
}

func NewAdmin(email, pass string, user *User) *Admin {
	return &Admin{
		email,
		pass,
		*user,
		Student{
			"Suraj",
		},
	}
}