package main

import "fmt"

func main() {
	// Mapas são tipados
	// Mapas precisam ser inicializados
	// Não aceita chaves repetidas
	aprovados := make(map[int]string)

	// Adicionando elementos ao mapa
	aprovados[1234] = "Maria"
	aprovados[5678] = "Pedro"
	aprovados[9874] = "Ana"
	fmt.Println(aprovados)

	// Iterando sobre o mapa usando range
	for key, value := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", value, key)
	}

	// Deletando um elemento do mapa
	delete(aprovados, 5678)
	fmt.Println(aprovados)
}
