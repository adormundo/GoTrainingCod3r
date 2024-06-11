package main

import (
	"fmt"
	"time"
)

// isPrimo verifica se um número é primo
func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// primos envia os primeiros números primos para um canal
func primos(n int, c chan int) {
	defer close(c) // Garantir que o canal será fechado ao final da função

	for i := 2; i <= n; i++ {
		if isPrimo(i) {
			c <- i
			time.Sleep(time.Millisecond * 100) // Adicionado um pequeno delay para simular processamento
		}
	}
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("\nFim!")
}
