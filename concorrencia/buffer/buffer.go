package main

import "fmt"

// rotina envia 6 valores inteiros para o canal ch
func rotina(ch chan int) {
	ch <- 1                 // Envia o valor 1 para o canal ch
	ch <- 2                 // Envia o valor 2 para o canal ch
	ch <- 3                 // Envia o valor 3 para o canal ch
	ch <- 4                 // Envia o valor 4 para o canal ch
	ch <- 5                 // Envia o valor 5 para o canal ch
	ch <- 6                 // Envia o valor 6 para o canal ch
	fmt.Println("Executou") // Imprime "Executou" após enviar todos os valores
}

func main() {
	// Cria um canal de inteiros com capacidade para 3 valores
	ch := make(chan int, 3)

	// Inicia uma goroutine para executar a função rotina, passando o canal ch
	go rotina(ch)

	// Recebe e imprime o primeiro valor enviado pelo canal ch
	fmt.Println(<-ch)
}
