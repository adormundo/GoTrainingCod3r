package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // Cria um slice s1 com 10 elementos e capacidade inicial de 20
	s2 := append(s1, 1, 2, 3) // Cria um novo slice s2 com base no s1 e adiciona mais elementos

	fmt.Println("Após a criação:")
	fmt.Println("s1:", s1) // s1 tem 10 elementos inicializados com 0
	fmt.Println("s2:", s2) // s2 é uma cópia de s1 com 13 elementos (10 + 1 + 2 + 3)

	s1[0] = 7 // Modifica o primeiro elemento de s1 para 7
	s1[9] = 9 // Modifica o último elemento de s1 para 9

	fmt.Println("\nApós as modificações em s1:")
	fmt.Println("s1:", s1) // s1 foi modificado
	fmt.Println("s2:", s2) // s2 também foi modificado, pois compartilha o mesmo array subjacente

	fmt.Println("\nTamanhos e capacidades:")
	fmt.Println("len(s1):", len(s1)) // Comprimento de s1 é 10
	fmt.Println("cap(s1):", cap(s1)) // Capacidade de s1 é 20
	fmt.Println("len(s2):", len(s2)) // Comprimento de s2 é 13
	fmt.Println("cap(s2):", cap(s2)) // Capacidade de s2 é 20 (mesma do s1)
}
