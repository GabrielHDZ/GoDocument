package main

import "fmt"

func main() {
	//inmutables enteros
	var arr1 [5]int
	arr1[1] = 26
	fmt.Println(arr1)
	//declaracion simple
	arr2 := [5]int{23, 12, 12, 34, 76}
	fmt.Println(arr2)
	//declaracion explicita
	arr3 := [...]int{12, 23, 34, 45, 45, 56, 67, 45}
	fmt.Println(arr3)
	//asignacion de valor explicito
	// posicion 2,3 y 4 quedan vacias, pero existentes
	arr4 := [...]string{0: "valor0", 1: "valor1", 5: "valortres"}
	fmt.Println(arr4)
	//se obtiene un array a partir de otro, termino "Slice"
	subArr := arr4[0:2]
	fmt.Print(subArr)
}
