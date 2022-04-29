package main

import "fmt"

func main() {
	var a1 [5]string
	a1[0] = "posicao 1"
	fmt.Println(a1)

	a2 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(a2)

	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a3)

	s1 := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(s1)

	s2 := a2[1:3] // inclusivo : exclusivo
	fmt.Println(s2)

	a2[1] = "Alterado"
	fmt.Println(s2)

	s3 := make([]float32, 10, 15)
	fmt.Println(s3)
	fmt.Println(len(s3)) // tamanho
	fmt.Println(cap(s3)) // capacidade
}
