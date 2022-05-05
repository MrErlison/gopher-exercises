package main

import "fmt"

func soma(numeros ...int) int {
	totalSoma := 0
	for _, numero := range numeros {
		totalSoma += numero
	}
	return totalSoma
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(total)

	escrever("Texto", 10, 20, 30, 40)
}
