package main

func main() {
	numero := 5
	cambioPorValor(numero, 15)
	cambioPorRefencia(&numero, 12)

}
func cambioPorValor(numero int, valor int) {
	numero = valor

}
func cambioPorRefencia(numero *int, valor int) {
	*numero = valor
}
