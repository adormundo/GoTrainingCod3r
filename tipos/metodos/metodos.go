package main

import (
	"fmt"
	"strings"
)

// Definição da struct pessoa
type pessoa struct {
	nome      string
	sobrenome string
}

// Método getNomeCompleto retorna o nome completo da pessoa
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Método setNomeCompleto atualiza o nome e sobrenome da pessoa
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	// Criando uma instância da struct pessoa
	p1 := pessoa{"Alvaro", "Dormundo"}

	// Chamando o método getNomeCompleto()
	fmt.Println(p1.getNomeCompleto()) // Saída: Alvaro Dormundo

	// Chamando o método setNomeCompleto()
	p1.setNomeCompleto("Jirlandia Jesus")

	// Chamando novamente o método getNomeCompleto() para verificar as mudanças
	fmt.Println(p1.getNomeCompleto()) // Saída: Jirlandia Jesus
}
