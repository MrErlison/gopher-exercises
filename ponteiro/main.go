package main

import "fmt"

func invert(n int) int {
	return n * -1
}

func invertnum(num *int) {
	*num *= -1
}

func main() {
	numero := 20
	numeroinvertido := invert(numero)
	fmt.Println(numero, numeroinvertido)

	invertnum(&numero)
	fmt.Println(numero)

}
