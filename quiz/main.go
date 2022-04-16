package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problema struct {
	pergunta string
	resposta string
}

func parseLinhas(linhas [][]string) []problema {

	problemas := make([]problema, len(linhas))

	for i, linha := range linhas {
		problemas[i] = problema{
			pergunta: linha[0],
			resposta: strings.TrimSpace(linha[1]),
		}
	}

	return problemas
}

func main() {

	csvArquivo := flag.String("csv", "problems.csv", "um arquivo cvs no formato 'pergunta,resposta'")
	flag.Parse()

	arquivo, err := os.Open(*csvArquivo)
	if err != nil {
		fmt.Println("Erro: ", err)
		os.Exit(1)
	}

	r := csv.NewReader(arquivo)
	linhas, err := r.ReadAll()
	if err != nil {
		fmt.Println("Erro: ", err)
		os.Exit(1)
	}

	problemas := parseLinhas(linhas)

	var resposta string
	acertos := 0
	for i, problema := range problemas {
		fmt.Printf("Problema %d: %s = ", i+1, problema.pergunta)
		fmt.Scanf("%s\n", &resposta)
		if resposta == problema.resposta {
			acertos++
		}
	}
	fmt.Printf("Acertou %d de %d\n", acertos, len(problemas))
}
