package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hola mundo")
	/**
	declaracion de variables
	*/
	var objeto string = "cadena de texto"
	var objeto2 int = 25
	//otra declaracion
	objeto3 := "cadena de texto2"
	fmt.Println(objeto)
	fmt.Println(objeto2)
	fmt.Println(objeto3)
	var nombre, app1, app2 = "carlos", "cardenas", "castillo"
	edad, peso := 45, 11
	fmt.Println(nombre, app1, app2, edad, peso)
	const constante1 = "cadena de texto constante, eso no se puede modificar mas adelante"

	//operadores
	/**
	>
	<
	>=
	<=
	==
	!=
	*/
	var edad2 = 10
	var prueba = edad2 > 10 //prueba obtiene un valor booleano
	fmt.Println(prueba)

	var constante2 = constante1 == "cadena de texto contante"
	fmt.Println(constante2)
	//condicionales
	// && || !
	ternario := 12 == 12 && "letra" == "letra" && (5 > 3 || 5 < 10)
	fmt.Println(ternario)

	const (
		lunes int = iota + 10 //por defecto inicia en 0 pero esta vez inica en 10
		martes
		miercoles
		jueves
		viernes
	)
	fmt.Println(jueves)

	texto3 := "cadena de texto para prueba"
	longitud := len(texto3) //longitud de la cadena de texto que se le pasa
	fmt.Println(longitud)
	caracter := texto3[3]                   //--> obtencion de codigo ASCII del caracter indicado, la cadena empieza en 0
	ultimoCaracter := texto3[len(texto3)-1] //obtebcion del ultimo caracter
	fmt.Println(caracter)                   //impresion del caracter en codigo ascii
	fmt.Println(reflect.TypeOf(caracter))   //impresion de tipo de caracter que es utf8
	fmt.Printf("%c\n", caracter)
	fmt.Print(ultimoCaracter) //tranformacion a letra legible
}
