package main

import "fmt"

func main() {
	//ciclo infinito
	cic := 1
	for {
		fmt.Println(cic)
		cic++
		if cic == 10 {
			//break o panic
			//panic("se cumple la condicion")
			break
		}
	}
	//ciclo for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//simulando un while
	//numero:=5
	contador := 0
	for contador < 10 {
		contador++
	}
	fmt.Println(contador)

	//foreach
	animales := []string{"perro", "gato", "oso", "elefante"}
	for indice, elemento := range animales {
		fmt.Println("indice del arreglo", indice, "animal", elemento)
	}
	//si no se quiere utilizar una variable declarada, sea indice o el elemento
	for _, elemento := range animales {
		fmt.Println("animal", elemento)

	}
	//for como do-while
	numeros := 10
	for isPass := true; isPass; isPass = numeros < 10 {
		fmt.Println(numeros)
		numeros++
	}
	//break y continue
	for i := 0; i <= 10; i++ {
		if i == 3 {
			fmt.Println("se ingresa a la acondicion, pero continuamos")
			continue
		}
		if i == 9 {
			fmt.Println("se ha enreado en la condicion y se termina todo")
			break
		}
		fmt.Println(i)
	}
}
