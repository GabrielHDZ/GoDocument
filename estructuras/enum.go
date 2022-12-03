package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {
	carlos := User{Username: "carlos", Type: Teacher}
	jose := User{Username: "Jose", Type: Student}

	fmt.Println(carlos)
	fmt.Println(jose)
}
