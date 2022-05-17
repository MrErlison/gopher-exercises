package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) imprime() {
	fmt.Println(u)
}

func main() {
	usuario01 := usuario{"Usuario", 18}
	usuario01.imprime()
}
