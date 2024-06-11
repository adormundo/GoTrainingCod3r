package main

import "fmt"

// multiplicacao retorna o produto de dois inteiros.
func multiplicacao(a, b int) int {
	return a * b
}

// exec recebe uma função que realiza uma operação em dois inteiros,
// além de dois inteiros como parâmetros, e retorna o resultado da operação.
func exec(operacao func(int, int) int, p1, p2 int) int {
	return operacao(p1, p2)
}

func main() {
	// Chama a função exec, passando a função multiplicacao e os parâmetros 3 e 4.
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado) // Imprime o resultado da multiplicação
}
