package main

import "fmt"

// Estrutura homogênea (tipado)
// Estrutura indexada (cada posição tem um index começando do 0)
// Estrutura estática (tamanho fixo)

func main() {
	// Declaração de um array de float64 com tamanho fixo 3
	var notas [3]float64
	fmt.Println("Array de notas inicializado:", notas)

	// Atribuição de valores ao array
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	notas[2] = 10 // Modificando o valor da terceira posição do array
	fmt.Println("Array de notas após atribuição de valores:", notas)

	// Cálculo da média das notas
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Com base nas %d notas, a média é %.2f\n", len(notas), media)
}
