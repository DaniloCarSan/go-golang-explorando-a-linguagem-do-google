package main

import (
	"fmt"
	"time"

	"github.com/DaniloCarSan/go_get_title_in_web_sites"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := go_get_title_in_web_sites.Titulo(url1)
	c2 := go_get_title_in_web_sites.Titulo(url2)
	c3 := go_get_title_in_web_sites.Titulo(url2)

	//estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam !"
		// default:
		// 	return "Sem resposta"
	}
}

func main() {

	campeao := oMaisRapido(
		"https://www.udemy.com/",
		"https://mail.google.com/",
		"https://wakatime.com/",
	)

	fmt.Println(campeao)

}
