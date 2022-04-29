package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 5
	multiplicacao := 2 * 10
	resto := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)

	v1, f1 := true, false
	fmt.Println(v1 && f1)
	fmt.Println(v1 || f1)
	fmt.Println(v1, f1)

	// ++ -- += -= *= /= %=
	n1 := 10
	fmt.Println(n1)
	n1++
	fmt.Println(n1)
	n1 += 10
	fmt.Println(n1)

}
