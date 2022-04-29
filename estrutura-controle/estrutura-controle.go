package main

import "fmt"

func main() {

	n1 := 10

	// > >= < <= != ==
	if n1 > 10 {
		fmt.Println("maior")
	} else if n1 < 10 {
		fmt.Println("menor")
	} else {
		fmt.Println("quaquer outra coisa")
	}

	if n2 := n1; n2 > 0 {
		fmt.Println("maior")
	} // a var n2 existe dentro do escopo {}

}
