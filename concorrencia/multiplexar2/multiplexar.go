package main

import (
	"fmt"
	"time"
)

// falar simula uma rotina que envia mensagens para um canal
func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c) // Fechar o canal ao terminar todas as goroutines
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return c
}

// juntar multiplexa (combina) mensagens de múltiplos canais em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c) // Fechar o canal ao terminar a função anônima
		for {
			select {
			case s, ok := <-entrada1:
				if !ok {
					entrada1 = nil // Marcar entrada1 como nil se for fechado
					continue
				}
				c <- s
			case s, ok := <-entrada2:
				if !ok {
					entrada2 = nil // Marcar entrada2 como nil se for fechado
					continue
				}
				c <- s
			}
			if entrada1 == nil && entrada2 == nil {
				break // Sai do loop se ambos os canais forem fechados
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))

	// Receber e exibir mensagens de c
	for msg := range c {
		fmt.Println(msg)
	}
}
