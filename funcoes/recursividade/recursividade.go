package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número invalido: %d", n)
	case n == 1:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {

	fat, _ := fatorial(12)

	fmt.Println(fat)

	_, err := fatorial(-1)
	if err != nil {
		fmt.Println(err)
	}

}
