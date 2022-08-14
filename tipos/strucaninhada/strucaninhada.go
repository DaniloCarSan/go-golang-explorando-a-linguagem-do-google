package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() (total float64) {

	total = 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return
}

func main() {

	p := pedido{
		userID: 1,
		itens: []item{
			{1, 4, 6.45},
			{2, 2, 23.60},
			{3, 5, 45.77},
			{3, 15, 1.70},
			{
				produtoID: 12,
				qtde:      23,
				preco:     0.56,
			},
		},
	}

	fmt.Printf("O valor total do pedido é : %.2f", p.valorTotal())

}
