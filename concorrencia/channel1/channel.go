package main

import "fmt"

func main() {
	channel := make(chan int, 1)

	channel <- 1 // enviando dados para o canal (escrita/write)
	<-channel    // recebendo dados do canal (leitura/reading)

	channel <- 2 // tentando enviar o segundo valor para o canal

	fmt.Println(<-channel) // recebendo e imprimindo o segundo valor do canal
}
