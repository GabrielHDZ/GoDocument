package main

import "fmt"

type User struct {
	Name string
	Edad int
}
type Student struct {
	User User
	Dni  string
}

func (self *User) setName(name string) {
	self.Name = name
}

// It's a method private
func (self *User) getName() string {
	return self.Name
}

//It's a method public
func (self *User) GetName() string {
	return self.getName()
}

func main() {
	var user1 User //instancia
	user1.Name = "carlos"
	user1.Edad = 23

	user2 := User{Name: "Jose", Edad: 44}

	Name := "Franco"
	Edad := 65
	user3 := User{Name, Edad}

	fmt.Println("usuario1", user1)
	fmt.Println("usuario2", user2)
	fmt.Println("usuario3", user3)

	user2.setName("Eduardo")
	fmt.Println(user2.getName())

	student1 := Student{User: user1, Dni: "SEDD3321"}
	fmt.Println(student1)

}
