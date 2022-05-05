package main

import "fmt"

func funcao1() {
	fmt.Println("Exec func 1")
}

func funcao2() {
	fmt.Println("Exec func 2")
}

func main() {
	defer funcao1() //adiar a exec até último momento
	funcao2()
}
