package main

import "fmt"

type Operacion func(valor1, valor2 int) int

func Operaciones(tipo string) Operacion {
	if tipo == "suma" {
		return func(valor1, valor2 int) int { return valor1 + valor2 }
	} else if tipo == "resta" {
		return func(valor1, valor2 int) int { return valor1 - valor2 }
	} else {
		return func(valor1, valor2 int) int { return valor1 * valor2 }
	}
}
func main() {
	operacion := Operaciones("resta")
	resultado := operacion(25, 50)
	fmt.Println(resultado)

}
