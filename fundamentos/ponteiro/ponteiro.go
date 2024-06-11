package main

import "fmt"

// IncrementarPonteiro incrementa o valor de um ponteiro para int.
func IncrementarPonteiro(p *int) {
	if p == nil {
		fmt.Println("Erro: ponteiro nulo")
		return
	}
	*p++
}

func main() {
	i := 1

	// Inicializando um ponteiro e Atribuindo o endereço da variável i ao ponteiro
	var ponteiro *int = &i

	// Incrementando o valor apontado pelo ponteiro
	IncrementarPonteiro(ponteiro)

	// Tentando incrementar um ponteiro nulo
	IncrementarPonteiro(nil)

	// Imprimindo o valor incrementado
	fmt.Println(*ponteiro) // Saída: 2
}
