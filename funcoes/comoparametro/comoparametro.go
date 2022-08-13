package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(multiplicacao func(int, int) int, p1, p2 int) int {
	return multiplicacao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 2, 5)
	fmt.Println(resultado)
}
