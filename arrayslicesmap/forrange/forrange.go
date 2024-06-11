package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Criando array sem passar o tamanho [...]
	numeros := [...]int{1, 2, 3, 4, 5} // Compilador conta!
	fmt.Println(numeros, reflect.TypeOf(numeros))

	// Iterando com range usando o índice
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// Iterando com range sem utilizar o índice
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
