package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos (composiçao)
	turboLigado bool
}

func main() {

	f := ferrari{}
	f.nome = "F30"
	f.turboLigado = true
	f.velocidadeAtual = 80

	fmt.Println(f)
	fmt.Println(f.turboLigado)

}
