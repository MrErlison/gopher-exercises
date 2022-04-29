package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	faculdade string
	curso     string
}

func main() {
	p1 := pessoa{"Nome", "Sobrenome", 20, 182}
	fmt.Println(p1)

	e1 := estudante{p1, "Faculdade", "Curso"}
	fmt.Println(e1)
	fmt.Println(e1.nome, e1.altura)
}
