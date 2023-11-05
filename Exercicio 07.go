package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Ricardo Souza Moraes"

	contarpalavras := criarmapa(frase)

	for palavra, contarletra := range contarpalavras {
		fmt.Print("Palavra: ", palavra)
		for letra, contar := range contarletra {
			fmt.Printf("%c: %d\n", letra, contar)
		}
	}
}

func contarletras(s string) map[rune]int {
	letras := make(map[rune]int)

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			letras[char]++
		}
	}
	return letras
}

func criarmapa(f string) map[string]map[rune]int {
	palavras := strings.Fields(f)
	contarpalavras := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		contarpalavras[palavra] = contarletras(palavra)
	}
	return contarpalavras
}
