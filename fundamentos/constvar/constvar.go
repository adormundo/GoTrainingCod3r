package main

import (
	"fmt"
	m "math" // Adicionando apelido (m) para a biblioteca math
)

// Função para calcular a área da circunferência
func calculateCircleArea(raio float64) float64 {
	const PI float64 = 3.1415
	return PI * m.Pow(raio, 2)
}

// Função principal onde o programa inicia
func main() {

	// Cálculo da área da circunferência
	raio := 3.2
	area := calculateCircleArea(raio)
	fmt.Println("A área da circunferência é", area)

	// Declaração de constantes
	const (
		a = 1
		b = 2
	)

	// Declaração de variáveis
	var (
		c = 3
		d = 4
	)
	fmt.Println("Valores de constantes e variáveis:", a, b, c, d)

	// Declaração de variáveis booleanas
	var e, f bool = true, false
	fmt.Println("Valores booleanos:", e, f)

	// Declaração e inicialização usando shorthand
	g, h, i := 2, false, "epa!"
	fmt.Println("Valores de variáveis usando shorthand:", g, h, i)
}
