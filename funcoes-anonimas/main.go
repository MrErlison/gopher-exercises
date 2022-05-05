package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Olá") // () indica que deve executar

	retorno := func(texto string) string {
		return texto
	}("Olá") // () indica que deve executar

	fmt.Println(retorno)
}
