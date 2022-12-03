package main

import "fmt"

func main() {
	//un slice es una nreferencia de uno o una parte de un arreglo
	numeros := []int{1, 2, 3, 4, 5, 6}

	numeros = append(numeros, 7)
	numeros = append(numeros, 8)
	sliceUno := numeros[0:3]
	sliceDos := numeros[2:5]
	numeros[2] = 503
	fmt.Println(numeros)
	fmt.Println(sliceUno)
	fmt.Println(sliceDos)

	//example
	meses := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio"}
	//propiedades de los slice
	/*
		puntero
		longitud y
		capaciidad
	*/
	longitud := len(meses)
	capacidad := cap(meses) //valor de capacidad maxima a 6 espacios maximos
	fmt.Printf("longitud del array %v, capacidad del mismo %v %p\n", longitud, capacidad, meses)

	meses = append(meses, "julio")
	meses = append(meses, "agosto")
	longitud = len(meses)
	capacidad = cap(meses) //valor de capacidad maxima aumento a 12 espacios maximos
	fmt.Printf("longitud del array %v, capacidad del mismo %v %p\n", longitud, capacidad, meses)

	//funcion make slace
	slice1 := make([]int, 3, 3) //en la construccion inicial del slace se indica su tipo,valores iniciales y la capacidad maxima exitdel mismo
	slice1[1] = 2               //se agrega el valor 2 a la posicion 1
	slice1 = append(slice1, 4)  //se agrega al slice un nuevo dato y tambien se duplica su capacidad maxima
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
