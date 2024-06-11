package main

import (
	"errors"
	"fmt"
)

// Fatorial calcula o fatorial de um número não negativo
func Fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return 0, errors.New("número inválido: deve ser não negativo")
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, err := Fatorial(n - 1)
		if err != nil {
			return 0, err
		}
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, err := Fatorial(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	_, err = Fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
