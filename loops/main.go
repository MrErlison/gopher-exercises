package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		//time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	for i = 100; i < 110; i++ {
		//time.Sleep(time.Second)
		fmt.Println(i)
	}

	nomes := [3]string{"nome1", "nome2", "nome3"}

	for indice, nome := range nomes {
		fmt.Println("Indice:", indice, " - nome:", nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println("indice:", indice, " - valor ascii:", letra, " - letra:", string(letra))
	}

	usuario := map[string]string{
		"nome":      "Nome Completo",
		"sobrenome": "Sobrenome",
	}
	for chave, valor := range usuario {
		fmt.Println("chave:", chave, " - valor:", valor)
	}

	// loop infinito
	// for { }
}
