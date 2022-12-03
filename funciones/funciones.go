package main

import "fmt"

func suma(num1, num2 int) (int, bool, string) {
	return num1 + num2, false, "hola"
}

//puntero o referencia de espacio en memoria
func ModificarReferenciasMemoria(propiedad1 *string) {
	//se cambia directamente el valor de la referencia en el espacio de memoria
	*propiedad1 = "nuevo valor de cadena de texto"
	//espacio en memoria
	fmt.Printf("%p\n", &propiedad1)

}

type TipoFactorial func(numero int) int

// It calculates the factorial of a number.
func Factorial(numero int) int {
	if numero == 1 {
		return 1
	}
	return Factorial(numero-1) * numero
}

func main() {
	cadena := "valor de cadena de texto"
	fmt.Printf("%p\n", &cadena)
	//el valor de cadena es cambiado en la funcion
	ModificarReferenciasMemoria(&cadena)

	fmt.Println("cadena >>>", cadena)
	//Recursividad
	var factorial TipoFactorial
	factorial = Factorial
	factor := factorial(5)
	fmt.Println("el factorial >>>", factor)

	//guion bajo, para omitir el segundo valor recibido
	resultado, _, _ := suma(5, 2)
	fmt.Println(resultado)

	//funcion annonima
	func() {
		fmt.Println("f anonima")
	}()

	resta := func(num1, num2 int) (int, bool) {
		return num1 - num2, false
	}

	restaFinal, _ := resta(5, 2)
	fmt.Println(restaFinal)
}
