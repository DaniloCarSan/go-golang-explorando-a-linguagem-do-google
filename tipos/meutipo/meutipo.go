package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(8, 10.0) {
		return "A"
	} else if n.entre(6, 7.99) {
		return "B"
	} else if n.entre(4, 5.99) {
		return "C"
	}

	return "D"
}

func main() {

	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(1))

}
