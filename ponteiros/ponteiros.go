package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)
	v2++
	fmt.Println(v1, v2)

	//ponteiros
	var v3 int
	var p1 *int

	v3 = 100
	p1 = &v3
	fmt.Println(v3, *p1, p1)
	v3++
	fmt.Println(v3, *p1, p1)
}
