package main

import (
	"fmt"
	"time"
)

func main() {
	// Comparação de strings
	a, b := "Banana", "Banana"
	fmt.Println("Strings iguais: ", a == b) // true

	// Comparação de inteiros
	fmt.Println("3 != 2:", 3 != 2) // true
	fmt.Println("3 < 2:", 3 < 2)   // false
	fmt.Println("3 > 2:", 3 > 2)   // true
	fmt.Println("3 <= 2:", 3 <= 2) // false
	fmt.Println("3 >= 2:", 3 >= 2) // true

	// Comparação de datas
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas iguais (==):", d1 == d2)        // true
	fmt.Println("Datas iguais (Equal):", d1.Equal(d2)) // true

	// Definição do tipo Pessoa
	type Pessoa struct {
		Nome  string
		Idade int
	}

	// Comparação de structs
	p1 := Pessoa{"João", 12}
	p2 := Pessoa{"João", 15}
	fmt.Println("Pessoas iguais:", p1 == p2) // false
}
