package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são impementadas implicitamente (automatico pelo compilador)
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Danilo", "Santos"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Livro", 289.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	imprimir(produto{"Celular", 29.90})

	fmt.Println(produto{"Chocolate", 9.90}.toString())

}
