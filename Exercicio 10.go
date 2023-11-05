package main

import "fmt"

func countPairsFrequency(numbers []int) map[[2]int]int {
	pairFrequency := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairFrequency[pair]++
		}
	}
	return pairFrequency
}

func main() {
	numbers := []int{1, 2, 3, 1, 2, 3, 4, 5}
	pairFrequency := countPairsFrequency(numbers)

	fmt.Println("Frequência de pares de números:")
	for pair, frequency := range pairFrequency {
		fmt.Printf("[%d, %d]: %d\n", pair[0], pair[1], frequency)
	}
}
