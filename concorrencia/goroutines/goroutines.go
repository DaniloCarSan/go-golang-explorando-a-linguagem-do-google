package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {

	// fmt.Printf("-----------------SEQUENCIAL---------------\n")
	// fale("Maria", "Porque você não fala comigo ?", 3)
	// fale("Maria", "Por que só posso falar depois de você !", 1)
	// fmt.Printf("-----------------FIM SEQUENCIAL---------------\n")

	// go fale("Maria", "Ei....", 500)
	// go fale("João", "Opa...", 500)
	// time.Sleep(time.Second + 5)
	// fmt.Printf("FIM@\n")

	go fale("Maria", "Entendi...", 10)
	fale("João", "Parabéns !", 5)
	fmt.Printf("FIM@\n")

}
