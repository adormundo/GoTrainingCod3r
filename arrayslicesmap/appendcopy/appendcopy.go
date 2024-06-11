package main

import "fmt"

func main() {
	// Inicializando slices
	var slice1 = []int{}     // slice1 com tamanho inicial 0 e capacidade 0
	slice2 := make([]int, 2) // slice2 com tamanho inicial 2 e capacidade 2

	// Exibindo estado inicial
	fmt.Println("Antes:")
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	// Adicionando elementos ao slice1
	slice1 = append(slice1, 4, 5, 6)

	// Exibindo estado após a adição de elementos
	fmt.Println("\nDepois de adicionar elementos a slice1:")
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	// Copiando elementos de slice1 para slice2
	copy(slice2, slice1) // copia de slice1 para slice2

	// Exibindo estado após a cópia
	fmt.Println("\nDepois de copiar slice1 para slice2:")
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))
}
