package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Code  int      `json:"code"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {

	// struct to json
	p1 := product{
		Code:  1,
		Name:  "LG GT-S2",
		Price: 8999.99,
		Tags:  []string{"Promoção", "Eletrônico", "Celular"},
	}
	// Primero valor o json e o segundo um erro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(p1Json)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 product
	jsonString := `{"code":2,"name":"Caneta","price":1.00,"tags":["Promoção","Importado","Bic"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])

}
