package main

import (
	"fmt"
	"math"
)

// Função para realizar operações aritméticas básicas
func basicArithmetic(a, b int) {
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)
}

// Função para realizar operações bitwise
func bitwiseOperations(a, b int) {
	fmt.Println("E (AND) =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou (OR) =>", a|b)   // 11 | 10 = 11
	fmt.Println("Xor (XOR) =>", a^b) // 11 ^ 10 = 01
}

// Função para utilizar operações da biblioteca Math
func mathOperations(c, d float64) {
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação (Potenciação) =>", math.Pow(c, d))
}

func main() {
	a := 3
	b := 2

	// Realizando operações aritméticas básicas
	basicArithmetic(a, b)

	// Realizando operações bitwise
	bitwiseOperations(a, b)

	// Utilizando a biblioteca Math para operações com ponto flutuante
	c := 3.0
	d := 2.0
	mathOperations(c, d)
}
