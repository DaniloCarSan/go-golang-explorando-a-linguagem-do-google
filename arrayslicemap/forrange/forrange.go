package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta e define o tamanho do array

	for index, value := range numeros {
		fmt.Printf("%d) %d \n", index, value)
	}

	for _, val := range numeros {
		fmt.Printf("valor: %d \n", val)
	}
}
