package main

import "fmt"

func main() {
	contas := map[string]float64{
		"Ricardo":  20.0,
		"Henrique": 130.0,
		"Felipe":   80.0,
	}
	saldos := calcularvalor(contas)

	fmt.Println("Saldo:")
	for pessoa, saldo := range saldos {
		fmt.Printf("%s: R$%.2f\n", pessoa, saldo)
	}
}

func calcularvalor(m map[string]float64) map[string]float64 {
	total := 0.0

	for _, valor := range m {
		total += valor
	}
	media := total / float64(len(m))

	saldos := make(map[string]float64)

	for pessoa, despesa := range m {
		saldo := despesa - media
		saldos[pessoa] = saldo
	}
	return saldos
}
