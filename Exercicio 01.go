package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Ricardo Souza Moraes"
	quantidade := ocorrencias(frase)

	for palavra, ocorrencia := range quantidade {
		fmt.Printf("%s: %d\n", palavra, ocorrencia)
	}

}

func ocorrencias(s string) map[string]int {
	palavras := strings.Fields(s)
	palavrascontadas := make(map[string]int)

	for _, palavra := range palavras {
		palavra = strings.ToLower(palavra)
		palavrascontadas[palavra]++
	}
	return palavrascontadas
}
