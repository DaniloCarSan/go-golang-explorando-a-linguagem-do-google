package main

import "fmt"

func main() {
	//Mapas devem ser inicializados
	//var aprovados map[int]string

	aprovados := make(map[int]string)
	aprovados[1] = "Maria"
	aprovados[2] = "João"
	aprovados[3] = "João"
	fmt.Println(aprovados)

	for key, value := range aprovados {
		fmt.Println(key, value)
	}

	fmt.Println(aprovados[3])

	delete(aprovados, 3)

	fmt.Println(aprovados)

}
