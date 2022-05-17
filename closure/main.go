package main

import "fmt"

func closure() func() {
	texto := "dentro de closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "dentro de main"
	fmt.Println(texto)

	funcaonova := closure()
	funcaonova()
}
