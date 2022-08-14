package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante
	fmt.Println("Só deipois que foi lido !")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c)

	fmt.Println("Foi lido")
	fmt.Println(<-c) ///deadlock !
	fmt.Println("Fim!!!")
}
