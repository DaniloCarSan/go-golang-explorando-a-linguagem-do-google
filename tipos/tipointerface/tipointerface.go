package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)
	fmt.Println(reflect.TypeOf(coisa))

	coisa = 3
	fmt.Println(coisa)
	fmt.Println(reflect.TypeOf(coisa))

	type dynamic interface{}

	var coisa2 dynamic = "Opa"
	fmt.Println(coisa2)
	fmt.Println(reflect.TypeOf(coisa2))

	coisa = true
	fmt.Println(coisa2)
	fmt.Println(reflect.TypeOf(coisa2))

}
