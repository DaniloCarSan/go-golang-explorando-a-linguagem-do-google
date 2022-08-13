package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
// (p produto): está disendo que está função está vinculada a struct produto
func (p produto) precoComDesconto() float64 {
	return p.preco - p.desconto
}

func main() {

	p1 := produto{
		nome:     "Bolo",
		preco:    123.69,
		desconto: 50.00,
	}

	fmt.Println(p1)
	fmt.Println(p1.precoComDesconto())

	p2 := produto{"Celular", 1234.67, 98.90}

	fmt.Println(p2)
	fmt.Println(p2.precoComDesconto())
}
