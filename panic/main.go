package main

import "fmt"

func media(valor1, valor2 float32) float32 {
	media := (valor1 + valor2) / 2

	if media > 6 {
		return media
	}

	panic("média abaixo de 6")

}

func main() {
	fmt.Println(media(6, 5))
}
