package main

import (
	"fmt"
	m "math" // m renomear o import funciona como ex: package as packageName
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferio pelo compilador

	//form reduzida
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 5
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "opa!"

	fmt.Println(g, h, i)

}
