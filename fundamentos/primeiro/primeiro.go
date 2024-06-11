// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports.
*/
import "fmt"

// A função main é a porta de entrada para um programa em Go.
// Dentro do mesmo pacote não pode haver duas funções main.
func main() {
	// Print imprime texto no console sem adicionar uma nova linha.
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
}
