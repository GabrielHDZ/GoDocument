package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
	Cazar()
}
type Mascota interface {
	Jugar()
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
func (self Perro) Jugar() {
	fmt.Println("el perro juega")
}

func ejecutarAnimal(animal Animal) {
	fmt.Println("Es un animal")
	animal.Comer()
	animal.Dormir()
	animal.Cazar()
}
func ejecutarMascota(animal Mascota) {
	fmt.Println("es una mascota")
	animal.Jugar()
}

func main() {
	yuya := Perro{Nombre: "yuya"}
	ejecutarAnimal(yuya)
	ejecutarMascota(yuya)
}
