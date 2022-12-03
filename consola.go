package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float32

	fmt.Print("ingresa el nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Print("ingresa la edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Print("ingresa la altura: ")
	fmt.Scanf("%f", &altura)
	//utilizar funcion panic
	if edad == 0 {
		panic("la edad no puede ser 0")
	}

	fmt.Printf("nombre %s con edad %d y con altura %.2f\n", nombre, edad, altura)

}
