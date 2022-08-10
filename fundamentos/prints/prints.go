package main

import "fmt"

func main() {

	fmt.Print("Mesma ")
	fmt.Print("Linha.")

	fmt.Println("Nova")
	fmt.Println("Linha.")

	x := 3.141516

	// fmt.Println(" VALOR É "+x)

	//converte para string
	xs := fmt.Sprint(x)

	fmt.Println("VALOR É " + xs)
	fmt.Println("valor é ", x)
	fmt.Printf("O valor é %0.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa!"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

	//generico para vários tipos
	fmt.Printf("\n %v %v %v %v %v", a, b, c, d)
}
