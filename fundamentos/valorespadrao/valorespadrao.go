package main

import "fmt"

/*
  Em Go, as variáveis não inicializadas são automaticamente atribuídas com 
  valores padrão:
  Numéricos (int, float64), o valor padrão é 0. 
  Bool, o valor padrão é false. 
  String, o valor padrão é uma string vazia "". 
  Ponteiros, o valor padrão é nil.
*/

func main() {
	// Declaração de variáveis e inicialização com valores padrões
	var a int     // Valor padrão: 0
	var b float64 // Valor padrão: 0
	var c bool    // Valor padrão: false
	var d string  // Valor padrão: ""

	// Declaração de ponteiros e inicialização com nil
	var e *int     // Valor padrão: nil
	var f *float64 // Valor padrão: nil
	var g *bool    // Valor padrão: nil
	var h *string  // Valor padrão: nil

	// Imprimindo os valores padrão das variáveis
	fmt.Printf("int: %v, float64: %v, bool: %v, string: %q \n", a, b, c, d)
	fmt.Printf("*int: %v, *float64: %v, *bool: %v, *string: %v \n", e, f, g, h)
}
