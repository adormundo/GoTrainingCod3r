package main

import (
	"fmt"
	"reflect"
)

// Fatorial calcula o fatorial de um número não negativo usando recursão
func Fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := Fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main() {
	resultado := Fatorial(5)
	fmt.Println(resultado, reflect.TypeOf(resultado))
}
