package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"Danilo": 5698.58,
		"Lorem":  8947.89,
		"Mario":  6800.90,
	}

	delete(salarios, "NÃO EXISTE")

	fmt.Println(salarios)

	salarios["Juliano"] = 3500.58

	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

}
