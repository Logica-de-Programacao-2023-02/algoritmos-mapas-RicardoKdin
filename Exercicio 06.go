package main

import "fmt"

func main() {
	m1 := map[string]int{"arroz": 1, "feijao": 2, "salada": 3, "macarrao": 4}
	m2 := map[string]int{"arroz": 2, "feijao": 2, "salada": 1, "macarrao": 3}

	mapa := []map[string]int{m1, m2}
	totalpalavra := somarcontagens(mapa)

	for palavra, contar := range totalpalavra {
		fmt.Printf("%s: %d\n", palavra, contar)
	}

}

func somarcontagens(c []map[string]int) map[string]int {
	total := make(map[string]int)

	for _, palavras := range c {
		for palavra, contar := range palavras {
			total[palavra] += contar
		}
	}
	return total
}
