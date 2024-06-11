package main

import (
	"fmt"
	"strconv"
)

// Função principal onde o programa inicia
func main() {
	// Divisão entre float64 e int convertendo y para float64
	divideFloatAndInt()

	// Truncar valor float64 para int
	truncateFloat()

	// Demonstração de conversões entre diferentes tipos
	conversions()

	// Conversão de string para boolean e uso condicional
	parseBoolean()
}

// Função para dividir um float64 por um int
func divideFloatAndInt() {
	x := 2.4
	y := 2
	fmt.Println("Resultado da divisão:", x/float64(y))
}

// Função para truncar um float64 para int
func truncateFloat() {
	nota := 6.9
	notaFinal := int(nota) // Truncando a casa decimal
	fmt.Println("Nota truncada:", notaFinal)
}

// Função para demonstrar conversões entre tipos
func conversions() {
	// Convertendo inteiro para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// Convertendo string para inteiro
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Erro ao converter string para int:", err)
	} else {
		fmt.Println("123 convertido para inteiro e subtraído por 122:", num-122)
	}
}

// Função para converter string para boolean e usar condicionalmente
func parseBoolean() {
	isTrue, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Erro ao converter string para boolean:", err)
	} else if isTrue {
		fmt.Println("Verdadeiro")
	}
}
