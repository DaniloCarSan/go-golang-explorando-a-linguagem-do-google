package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	medelo          string
	turboLigado     bool
	delocodadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{"F430", false, 100}
	carro1.ligarTurbo()

	// Quando se está trabalhando em nivel de interface se passa (&)
	// a referencia da memória no exemplo acima estamos trabalhando com literal
	var carro2 esportivo = &ferrari{"F30", false, 100}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

}
