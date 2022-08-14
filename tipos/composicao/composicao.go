package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmwS7 struct {
	fazendoBaliza bool
	turboLigado   bool
}

func (b *bmwS7) ligarTurbo() {
	b.turboLigado = true
	fmt.Println("Turbo ligado....")
}

func (b *bmwS7) fazerBaliza() {
	b.fazendoBaliza = true
	fmt.Println("Fazer baliza....")
}

func main() {

	var b esportivoLuxuoso = &bmwS7{
		fazendoBaliza: false,
		turboLigado:   false,
	}

	b.ligarTurbo()
	b.fazerBaliza()

	fmt.Println(b)

}
