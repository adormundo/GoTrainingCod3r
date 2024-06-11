package main

import "fmt"

/*
	Essa demonstração ilustra como os closures podem capturar e reter variáveis
	de seu escopo léxico, permitindo a criação de funções com comportamento
	personalizado baseado no contexto em que foram definidas.

	Um closure em Go é uma função que referencia variáveis fora de seu
	próprio escopo. No exemplo, a função anônima printX captura a variável x
	definida no escopo da função closure. Isso significa que mesmo depois de
	closure ter retornado, a função anônima ainda mantém uma referência à
	variável x e pode acessar e imprimir seu valor.
*/

// closure retorna uma função que imprime o valor de x
func closure() func() {
	x := 10
	// printX uma função anônima que captura a variável x
	printX := func() {
		fmt.Println(x)
	}
	return printX
}

func main() {
	// Variável x no escopo da função main
	x := 20
	fmt.Println(x) // Imprime 20

	// Chamamos a função closure e armazenamos a função retornada
	imprimeX := closure()

	// Executamos a função retornada, que imprime o valor de x capturado no closure
	imprimeX() // Imprime 10
}
