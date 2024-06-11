package main

import (
	"fmt"
	"time"
)

func rotina(channel chan int) {
	time.Sleep(time.Second) // Simula algum processamento demorado
	channel <- 1            // Envia o valor 1 para o canal (operação bloqueante)
	fmt.Println("Só depois da leitura")
}

func main() {
	channel := make(chan int) // Canal sem buffer

	go rotina(channel) // Inicia uma goroutine para a função rotina

	// Lê um valor do canal (operação bloqueante até que haja um valor no canal)
	fmt.Println(<-channel)
	fmt.Println("Foi lido")

	// Tentativa de ler do canal novamente sem que haja outro valor enviado ao canal
	// Isso causará um deadlock, pois não há mais nenhuma goroutine pronta para enviar dados
	fmt.Println(<-channel)
	fmt.Println("Finalmente") // Essa linha nunca será alcançada devido ao deadlock
}
