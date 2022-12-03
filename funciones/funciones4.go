package main

import (
	"errors"
	"fmt"
)

//variadic function
func DataVariante(datos ...int) int {
	var promedio, suma int
	for _, calificacion := range datos {
		suma = suma + calificacion
	}

	promedio = suma / len(datos)
	return promedio
}
func Errores(variable1, variable2 int) (int, error) {
	if variable1 == 0 {
		return 0, errors.New("generamos un error al recibir 0 como primer parametro")
	}
	return variable2 / variable1, nil
}

func main() {
	promedios := DataVariante(7, 8, 5, 9, 10)
	fmt.Println(promedios)
	if resultado, err := Errores(0, 2); err != nil {
		fmt.Print("existe error ", err)
	} else {
		fmt.Print("resultado ", resultado)
	}
}
