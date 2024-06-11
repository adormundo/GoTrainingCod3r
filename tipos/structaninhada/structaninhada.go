package main

import "fmt"

// item representa um item do pedido com seu produto, quantidade e preço.
type item struct {
	produtoID int     // ID do produto
	qtd       int     // quantidade do produto
	preco     float64 // preço unitário do produto
}

// pedido representa um pedido com um usuário e uma lista de itens.
type pedido struct {
	userID int    // ID do usuário
	itens  []item // lista de itens do pedido
}

// valorTotal calcula o valor total do pedido.
func (p pedido) valorTotal() float64 {
	var total float64
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	// Criando um pedido de exemplo
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtd: 2, preco: 12.10},
			{produtoID: 2, qtd: 1, preco: 23.49},
			{produtoID: 11, qtd: 100, preco: 3.13},
		},
	}

	// Imprimindo o valor total do pedido
	fmt.Printf("Valor total do pedido é R$%.2f", pedido.valorTotal())
}
