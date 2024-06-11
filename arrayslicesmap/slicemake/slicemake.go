package main

import "fmt"

func main() {
	// Criando um slice com 10 elementos inicializados com zero
	s := make([]int, 10)
	s[9] = 12 // Atribuindo o valor 12 ao último elemento do slice
	fmt.Println(s)

	// Criando um slice com 10 elementos inicializados com zero e capacidade interna de 20
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// Adicionando elementos ao slice usando append
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Append após atingir a capacidade máxima do array interno, dobrou a capacidade para 40
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
