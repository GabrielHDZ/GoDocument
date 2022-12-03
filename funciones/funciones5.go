package main

import (
	"errors"
	"fmt"
	"os"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("not possible number 0")
	} else {
		return dividendo / divisor, nil
	}
}

func main() {
	if resultado, err := division(10, 0); err != nil {
		fmt.Println(err)
		//panic(err)
	} else {
		fmt.Println(resultado)
	}

	// A function that is executed after the main function is executed.
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("ocurrio un error edurante el proceso de lectura")
		}
	}()
	//uso de defer
	if file, err := os.Open("lectusra.txt"); err != nil {
		panic(err)
	} else {
		defer func() { fmt.Println("Cerramos el archivo") }()
		contenido := make([]byte, 254)
		longitud, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[0:longitud])
		fmt.Println(contenidoArchivo)
	}

}
