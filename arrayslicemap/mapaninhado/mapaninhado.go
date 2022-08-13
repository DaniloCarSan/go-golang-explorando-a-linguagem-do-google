package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 2443.89,
			"Gil":      9475.45,
			"Gabiru":   4957.89,
		},
		"J": {
			"Jose": 4894.90,
			"João": 9585.98,
		},
		"P": {
			"Pedro": 33000.00,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")

	for letra, funcionarios := range funcsPorLetra {
		fmt.Println(letra, funcionarios)
	}
}
