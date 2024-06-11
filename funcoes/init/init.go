package main

import "fmt"

// init é uma função especial em Go que é chamada automaticamente antes da função main.
func init() {
	fmt.Println("Inicializando...") // Imprime "Inicializando..."
}

func main() {
	fmt.Println("Main...") // Imprime "Main..."
}
