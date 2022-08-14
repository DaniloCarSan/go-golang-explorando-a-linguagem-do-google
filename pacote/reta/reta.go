package main

import "math"

//Iniciando com letra maiúscula é PUBLICO (visibilidade dentro e fora do pacote)
//Iniciando com letra minúscula é PRIVADO  (dentro do pacote não do arquivo)

// Exemplo
// Ponto - gerá algo público
// ponto ou _Ponto - gerá algo privado

// Ponto representa uma coordenada no plano catesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distacia é reponsável por calcular a distância linear entre dois potos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
