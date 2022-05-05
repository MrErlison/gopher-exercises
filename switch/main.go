package main

import "fmt"

func main() {
	numeroSemana := 1

	switch numeroSemana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Número da semana inválido")
	}

	numeroSemana += 3

	switch {
	case numeroSemana == 1:
		fmt.Println("Domingo")
	case numeroSemana == 2:
		fmt.Println("Segunda")
	case numeroSemana == 3:
		fmt.Println("Terça")
	case numeroSemana == 4:
		fmt.Println("Quarta")
		fallthrough //Passa para o próximo case sem avaliar
	case numeroSemana == 5:
		fmt.Println("Quinta")
	case numeroSemana == 6:
		fmt.Println("Sexta")
	case numeroSemana == 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Número da semana inválido")
	}

}
