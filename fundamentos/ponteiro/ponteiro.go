package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i
	// Go não possui aritmética de ponteiros
	// p++

	fmt.Println(p)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	i++
	fmt.Println(i)
	fmt.Println(*p)
}
