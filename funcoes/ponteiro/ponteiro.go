package main

import "fmt"

func inc1(n int) int {
	n++
	return n
}

// um ponteiro é representado por *

func inc2(n *int) {
	//revisão: o * é usado para desreferenciar , ou seja,
	//ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	fmt.Println("ORIGINAL", n)

	fmt.Println("INCREMENTA LOCAL: ", inc1(n))

	fmt.Println("CONTINUA ORIGINAL", n)

	inc2(&n)

	fmt.Println("ALTERADO PELA INC2 PASSANDO A REFERENCIA:", n)

}
