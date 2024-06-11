package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":   1234.45,
		"João":   15456.78,
		"Pedro":  1200.0,
		"Álvaro": 135697.80,
	}

	// Tentativa de adicionar um novo funcionário com salário
	funcsESalarios["Carlos"] = 980.50

	// Tentativa de deletar um funcionário que não existe
	delete(funcsESalarios, "Corno")

	for nome, salario := range funcsESalarios {
		fmt.Printf("Nome: %s - Salário: %.2f\n", nome, salario)
	}
}
