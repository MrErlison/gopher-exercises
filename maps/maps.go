package main

import "fmt"

func main() {
	m1 := map[string]string{
		"nome":      "John",
		"sobrenome": "Doe",
	}

	fmt.Println(m1)
	fmt.Println(m1["nome"])

	m2 := map[int]string{
		1: "John",
		2: "Doe",
	}

	fmt.Println(m2)
	fmt.Println(m2[1])

	m3 := map[string]map[string]string{
		"aluno": {
			"nome":      "John",
			"sobrenome": "Doe",
		},
		"curso": {
			"campus": "Campus 1",
		},
	}

	fmt.Println(m3)

	delete(m3, "nome")
	fmt.Println(m3)

}
