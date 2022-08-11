package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2
	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {

	tv50, tv32, sorvete := comprar(true, true)

	fmt.Printf("Tv50: %t Tv32: %t Sorvete: %t Saudável: %t \n", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = comprar(false, true)

	fmt.Printf("Tv50: %t Tv32: %t Sorvete: %t Saudável: %t \n", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = comprar(true, false)

	fmt.Printf("Tv50: %t Tv32: %t Sorvete: %t Saudável: %t \n", tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = comprar(false, false)

	fmt.Printf("Tv50: %t Tv32: %t Sorvete: %t Saudável: %t \n", tv50, tv32, sorvete, !sorvete)

}
