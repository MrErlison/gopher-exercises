package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calcula(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtrai := n1 - n2
	return soma, subtrai
}

func main() {

	s1 := somar(10, 20)
	fmt.Println(s1)

	s2, s3 := calcula(10, 20)
	fmt.Println(s2, s3)

	_, s4 := calcula(10, 20)
	fmt.Println(s4)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	r1 := f("Exibe texto")
	fmt.Println(r1)

}
