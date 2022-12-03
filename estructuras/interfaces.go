package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
	Cazar()
}

type Perro struct {
	Nombre string
}

func (self Perro) Comer() {
	fmt.Println("el perro come")
}
func (self Perro) Dormir() {
	fmt.Println("el perro duerme")
}
func (self Perro) Cazar() {
	fmt.Println("el perro caza")
}

func ejecutarAccion(animal Animal) {
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}

func main() {
	yuya := Perro{Nombre: "yuya"}
	ejecutarAccion(yuya)
}
