package main

import (
	"fmt"
	"time"
)

/*
   Channel (canal) - é a forma de comunicação entre as goroutines,
   ele é síncrono, sendo um ponto de parada.
   Channel possui o tipo, sendo chan
*/

func doisTresQuatroVezes(base int, channel chan int) {
	for i := 2; i <= 4; i++ {
		time.Sleep(time.Second)
		channel <- i * base
		fmt.Println("Loop", i)
	}
	close(channel) // fechando o canal após as operações terem sido concluídas
}

func main() {
	channel := make(chan int)
	go doisTresQuatroVezes(2, channel)
	fmt.Println("Passei da goroutine")

	// Recebendo os dados do canal
	for num := range channel {
		fmt.Println("Recebi do canal:", num)
	}

	fmt.Println("Programa terminou")
}
