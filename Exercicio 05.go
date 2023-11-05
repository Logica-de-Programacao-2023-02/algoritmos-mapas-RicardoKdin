package main

import "fmt"

func main() {
	s := "Ricardo Souza Moraes"
	quantidade := quantidadecaracter(s)

	for f, i := range quantidade {
		fmt.Printf("%c: %d\n", f, i)
	}

}

func quantidadecaracter(s string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, caractere := range s {
		frequencia[caractere]++
	}
	return frequencia
}
