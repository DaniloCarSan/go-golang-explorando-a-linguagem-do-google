package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 1:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {

	fat := fatorial(4)

	fmt.Println(fat)

}
