package main

import "fmt"

func main() {
	// Print não adiciona uma nova linha ao final da saída
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Println adiciona uma nova linha ao final da saída
	fmt.Println(" Nova")
	fmt.Println("linha.")
	fmt.Println("fim.")

	// Declaração e inicialização de variáveis
	x := 3.141516

	// Não podemos concatenar strings com valores de outros tipos diretamente
	// fmt.Println("O valor de x é" + x) --> Não funciona

	// Podemos imprimir múltiplos valores com vírgulas
	fmt.Println("O valor de x é", x)

	// Printf permite formatar a string de saída
	fmt.Printf("O valor de x é %.2f", x) // %.2f formata x para 2 casas decimais

	// fmt.Sprint converte x para uma string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	// Declaração e inicialização de variáveis de diferentes tipos
	a := 1     // int
	b := 1.999 // float64
	c := false // bool
	d := "Opa" // string

	// Printf com especificadores de formato
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// %d - inteiro
	// %f - float64
	// %.1f - float64 com 1 casa decimal
	// %t - booleano
	// %s - string

	// %v representa o valor em seu formato padrão
	fmt.Printf("\n%v %v %v %v\n", a, b, c, d)
}
