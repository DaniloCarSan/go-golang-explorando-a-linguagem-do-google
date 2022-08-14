package main

import (
	"fmt"

	"github.com/DaniloCarSan/go_get_title_in_web_sites"
)

func encaminhar(origin <-chan string, destino chan string) {
	for {
		destino <- <-origin
	}
}

// multiplexar - misturar (mensagens) num canal unico
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		go_get_title_in_web_sites.Titulo("https://www.udemy.com/", "https://mail.google.com/"),
		go_get_title_in_web_sites.Titulo("https://wakatime.com/", "https://github.com/"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
