package main

import "fmt"

func main() {

	// int8 int16 int32 int64
	var n1 int8 = -1
	fmt.Println(n1)

	var n2 uint8 = 1
	fmt.Println(n2)

	// int32 = rune
	var n3 rune = 12345
	fmt.Println(n3)

	//byte = uint8
	var n4 byte = 123
	fmt.Println(n4)

	// float32 e float64
	var f1 float32 = 123.12
	fmt.Println(f1)

	var s1 string = "string"
	fmt.Println(s1)

	var b1 bool = true
	fmt.Println(b1)

	var e1 error
	fmt.Println(e1)

}
