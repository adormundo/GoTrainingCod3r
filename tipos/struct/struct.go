package main

import (
	"fmt"
	"reflect"
)

// Produto representa um produto com nome, preço e desconto.
type Produto struct {
	Nome     string  // nome do produto
	Preco    float64 // preço do produto
	Desconto float64 // desconto aplicado ao produto
}

// precoComDesconto retorna o preço do produto com o desconto aplicado.
func (p Produto) precoComDesconto() float64 {
	return p.Preco * (1 - p.Desconto)
}

func main() {

	// Criando o primeiro produto
	p1 := Produto{
		Nome:     "Lápis",
		Preco:    1.79,
		Desconto: 0.05,
	}

	// Imprimindo informações sobre o primeiro produto
	fmt.Println(p1)                    // exibe o produto
	fmt.Println(reflect.TypeOf(p1))    // exibe o tipo do produto
	fmt.Println(p1.precoComDesconto()) // exibe o preço com desconto

	// Criando o segundo produto
	p2 := Produto{"Notebook", 1989.90, 0.10}

	// Imprimindo informações sobre o segundo produto
	fmt.Println(p2)                    // exibe o produto
	fmt.Println(reflect.TypeOf(p2))    // exibe o tipo do produto
	fmt.Println(p2.precoComDesconto()) // exibe o preço com desconto
}
