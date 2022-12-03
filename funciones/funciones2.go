package main

import "fmt"

type Operacion func(balance, cantidad int) int
type Ejecutador func(funcion Operacion, valor1 int, valor2 int)

//funciones de type Operacion
func Suma(valor1, valor2 int) int {
	return valor1 + valor2
}
func Resta(valor1, valor2 int) int {
	return valor1 - valor2
}

func Ejecutar(funcion Operacion, valor1 int, valor2 int) {
	fmt.Println("Inicia")
	resultado := funcion(valor1, valor2)
	fmt.Println(resultado)
	fmt.Println("Termina")

}

func main() {
	Ejecutar(Suma, 34, 12)
}
