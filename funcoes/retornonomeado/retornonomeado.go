package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) /*retorna na ordem definida aqui*/ {
	primeiro = p1
	segundo = p2
	return //retono limpo
}

func main() {
	r1, r2 := trocar(1, 2)
	fmt.Println(r1, r2)
}
