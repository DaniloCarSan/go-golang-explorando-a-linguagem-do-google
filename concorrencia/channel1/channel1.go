package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //enviando dados para canal(escrita)
	<-ch    // recebendo dado do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)

}
