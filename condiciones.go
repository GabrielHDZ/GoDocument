package main

import "fmt"

func main() {
	var edad int
	fmt.Println("Ingresar edad")
	fmt.Scanf("%d", &edad)

	if edad >= 18 {
		//si se cumple una condicion
		fmt.Println("Mayor de edad")
	} else if edad < 18 || edad > 15 {
		//sino
		fmt.Println("menor de edad, puberto")
	} else {
		fmt.Println("menor normal")
	}
	//las siguientes variables solo sirven el el bloque de codigo
	if nombre, altura := "cala", 15; nombre == "cala" {
		fmt.Println("nombre", nombre, "altura", altura)
	} else {
		fmt.Println("otra condicion")
	}

	/*
		condicion switch

	*/
	var estacion string
	fmt.Println("##### Ingresar la estacion del año #####")
	fmt.Scanf("%s", &estacion)
	switch estacion {
	case "primavera":
		fmt.Println("Primavera.")
		fmt.Println("La primavera​ es una de las cuatro estaciones del año, sigue al invierno y precede al verano.")
	case "verano":
		fmt.Println("Verano")
		fmt.Println("El verano o estío es una de las cuatro estaciones de las zonas templadas. Es la más cálida de ellas. Sigue a la primavera y precede al otoño.")
	case "otono":
		fmt.Println("Otoño")
		fmt.Println("El otoño es una de las cuatro estaciones del año y una de las cuatro de las zonas templadas. Sigue al verano y precede al invierno. Astronómicamente, comienza con el equinoccio de otoño y termina con el solsticio de invierno.")
	case "invierno":
		fmt.Println("Invierno")
		fmt.Println("El invierno es una de las cuatro estaciones de las zonas templadas. Sigue al otoño y precede a la primavera. Esta estación se caracteriza por días más cortos, noches más largas y temperaturas más bajas a medida que nos alejamos de la línea ecuatorial.")
	default:
		fmt.Println("no se ingreso ninguna estacion")
	}

	/*
		condicion map

	*/
	//declaracion de variable map
	persona := make(map[string]string)
	persona["moral"] = "juan"
	//evaluar la existencia del indice buscado
	/*
		person, isExists := persona["moral"]
		if isExists {
			fmt.Println(person)
		} else {
			fmt.Println("no existe esta persona")
		}*/
	//evaluar en la misma sentencia

	if person, isExists := persona["moral"]; isExists {
		fmt.Println(person)
	} else {
		fmt.Println("no existe esta persona")
	}
}
