package main

import "fmt"

// Definição da interface imprimivel
type imprimivel interface {
	toString() string
}

// Definição da struct pessoa
type pessoa struct {
	nome      string
	sobrenome string
}

// Definição da struct produto
type produto struct {
	nome  string
	preco float64
}

// Método toString para pessoa
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

// Método toString para produto
func (p produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)
}

// Função imprimir que aceita qualquer imprimivel
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// Exemplo com pessoa
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// Exemplo com produto
	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
