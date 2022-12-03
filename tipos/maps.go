package main

import "fmt"

func main() {
	//map es una colecicon desordenanda de elementos
	dias := make(map[int]string) //creacionde map, clave int y valor string
	dias[0] = "lunes"
	dias[1] = "martes"
	dias[2] = "miercoles"
	dias[3] = "jueves"
	dias[4] = "viernes"
	dias[5] = "sabado"
	dias[6] = "domingo"

	dias[1] = "MARTES"
	delete(dias, 2)

	fmt.Println(dias)

	//mapa de mapas
	calificaciones := make(map[string][]int)
	calificaciones["cad"] = []int{1, 2, 3, 4, 5, 67}
	fmt.Println(calificaciones)
	//recorrer el map y visualizar las diferentes impresiones que presenta
	for llave, valor := range dias {
		//la impresion no tiene un orden
		fmt.Println(llave, valor)
	}

}
