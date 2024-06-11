package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays e Slices em Go

	// Array
	a1 := [3]int{1, 2, 3}
	fmt.Println("Array a1:", a1, "Tipo:", reflect.TypeOf(a1))

	// Slice
	s1 := []int{1, 2, 3}
	fmt.Println("Slice s1:", s1, "Tipo:", reflect.TypeOf(s1))

	// Acessando um slice de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // Slice do array a2, do índice 1 ao 2 (3-1)
	fmt.Println("\nArray a2:", a2, "Tipo:", reflect.TypeOf(a2))
	fmt.Println("Slice s2:", s2, "Tipo:", reflect.TypeOf(s2))

	// Novo slice, apontando para o mesmo array a2
	s3 := a2[:2] // Slice do array a2, do início até o índice 1 (2-1)
	fmt.Println("Slice s3:", s3, "Tipo:", reflect.TypeOf(s3))

	// Criando um slice de um slice
	s4 := s2[:1] // Slice do slice s2, do início até o índice 0 (1-1)
	fmt.Println("Slice s4:", s4, "Tipo:", reflect.TypeOf(s4))
}
