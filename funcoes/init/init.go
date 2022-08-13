package main

import "fmt"

func init() {
	fmt.Println("Inicializando.....")
}

func main() {
	fmt.Println("Main.....")
}

// rodar no diretorio init: go run *.go
// sera imprimido
// Inicializando.....
// Inicializando 2.....
// Main.....
