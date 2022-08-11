package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("Som", a+b)
	fmt.Println("Sub", a-b)
	fmt.Println("Dev", a/b)
	fmt.Println("Mult", a*b)
	fmt.Println("Mod", a%b)

	//bitwise
	fmt.Println("E => ", a&b)  // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 10
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//math
	fmt.Println("Max =>", math.Max(float64(a), float64(b)))
	fmt.Println("Min =>", math.Min(float64(c), float64(d)))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
	fmt.Println("Ceil =>", math.Ceil(23.45))

}
