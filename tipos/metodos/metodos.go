package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomecompleto string) {
	partes := strings.Split(nomecompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Mario", "Cart"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Danilo Carreiro")

	fmt.Println(p1.getNomeCompleto())
}
