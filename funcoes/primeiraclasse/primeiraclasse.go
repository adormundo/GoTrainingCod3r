package main

import "fmt"

// Soma é uma função anônima (closure) que retorna a soma de dois inteiros
var Soma = func(a, b int) int {
	return a + b
}

func main() {
	// Chamando a função anônima Soma diretamente
	fmt.Println(Soma(2, 3))

	// Declarando e atribuindo uma função anônima (closure) à variável sub
	sub := func(a, b int) int {
		return a - b
	}

	// Chamando a função anônima sub
	fmt.Println(sub(2, 3))
}
