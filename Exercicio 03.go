package main

import "fmt"

func main() {
	numeros := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	total := somatudo(numeros)
	fmt.Print("Soma total: ", total)

}

func somatudo(m map[string]int) int {
	soma := 0
	for _, valor := range m {
		soma += valor
	}
	return soma
}
