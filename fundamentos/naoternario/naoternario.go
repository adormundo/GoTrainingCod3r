package main

import "fmt"

// Dominio: Lógica de Negócio

// ObterResultado calcula o resultado baseado na nota.
func ObterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

// Infraestrutura: Entrada/Saída

// ImprimirResultado imprime o resultado no console.
func ImprimirResultado(nota float64) {
	resultado := ObterResultado(nota)
	fmt.Println(resultado)
}

// Entrada do Programa

func main() {
	nota := 6.2
	ImprimirResultado(nota)
}
