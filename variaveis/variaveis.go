package main

import "fmt"

func main() {

	var v1 string = "v1"
	fmt.Println(v1)

	v2 := "v2"
	fmt.Println(v2)

	var (
		v3 string = "variavel3"
		v4 string = "variavel4"
	)

	fmt.Println(v3, v4)

	v5, v6 := "v5", "v6"
	fmt.Println(v5, v6)

}
