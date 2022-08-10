package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//tipos numericos
	fmt.Println(1, 2, 100)
	fmt.Println("Litera interio é do tipo", reflect.TypeOf(32000))

	//sem sinal (só positivos) ... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("o byte é ", reflect.TypeOf(b))

	//com sinal  ... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O Valor máximo do int é", i1)
	fmt.Println("O tipo i1", reflect.TypeOf(i1))

	var i2 rune = 'a' //representação um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais (float32,float64)
	var x float32 = 49.99
	fmt.Println("O tipo é", reflect.TypeOf(x))

}
