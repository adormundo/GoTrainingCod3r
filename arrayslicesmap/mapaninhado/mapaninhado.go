package main

import "fmt"

type Funcionario struct {
	Nome    string
	Salario float64
}

func main() {
	funcionariosPorLetra := map[string]map[string]Funcionario{
		"G": {
			"Gabriela Silva": {Nome: "Gabriela Silva", Salario: 15456.78},
			"Guga Junior":    {Nome: "Guga Junior", Salario: 8456.78},
		},
		"J": {
			"José João": {Nome: "José João", Salario: 11325.45},
		},
		"A": {
			"Amostradin": {Nome: "Amostradin", Salario: 99999.99},
		},
	}

	// Deletando o mapa associado à chave "G"
	delete(funcionariosPorLetra, "G")
	fmt.Println(funcionariosPorLetra)

	// Iterando sobre o mapa após deletar
	for letra, funcionarios := range funcionariosPorLetra {
		for nome, funcionario := range funcionarios {
			fmt.Printf("Letra: %s - Nome: %s, Salário: %.2f\n", letra, nome, funcionario.Salario)
		}
	}
}
