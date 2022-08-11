package main

import "fmt"

// não há operador ternario
// nota >= 6 ? "Aprovado" : "Reprovado"
func resultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(resultado(8))
}
