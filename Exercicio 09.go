package main

import "fmt"

func generateFibonacci(n int) map[int]int {
	fibonacciMap := make(map[int]int)

	a, b := 0, 1
	index := 0

	for a <= n {
		fibonacciMap[index] = a
		a, b = b, a+b
		index++
	}

	return fibonacciMap
}

func main() {
	n := 100
	fibonacciSequence := generateFibonacci(n)

	fmt.Println("Sequência de Fibonacci até", n, ":")
	for index, value := range fibonacciSequence {
		fmt.Printf("Fib[%d]: %d\n", index, value)
	}
}
