package main

import "fmt"

type usuario struct {
	nome     string
	idade    string
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {

	var u1 usuario

	e1 := endereco{"logradouro", 10}

	u1.nome = "Usuario 1"
	u1.idade = "18"
	fmt.Println(u1)

	u2 := usuario{"Usuario 2", "21", e1}
	fmt.Println(u2)

	u3 := usuario{nome: "Usuario 3"}
	fmt.Println(u3)

}
