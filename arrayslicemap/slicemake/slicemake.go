package main

import "fmt"

func main() {
	// Ele cria uma array interno
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Cria uma slice de 10 posições e um array interno de 20 posiçoes
	s = make([]int, 10, 20)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
