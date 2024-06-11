package main

import (
	"fmt"
	"time"
)

func main() {
	// Usando for como um while (não existe while em Go)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// For com inicialização, condição e pós-execução
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println() // Nova linha após a sequência

	// Misturando repetição com controle condicional
	fmt.Println("Misturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Par")
		} else {
			fmt.Println(i, "Ímpar")
		}
	}

	// Laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
		// Adicione uma condição de escape se necessário para evitar um loop infinito durante o desenvolvimento
	}
}
