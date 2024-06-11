package main

import (
	"fmt"
	"time"
)

// Função fale simula uma pessoa falando várias vezes com intervalos de um segundo
func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second) // Aguarda um segundo
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Exemplo de chamada de função normal (sem goroutine)
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// Exemplo de chamada de função utilizando goroutine
	go fale("Maria", "Ei...", 10) // Inicia uma goroutine para Maria falar 10 vezes
	fale("João", "Opa...", 5)     // João fala 5 vezes na thread principal

	// Como o programa continua imprimindo na tela, ele faz com que o processamento do código continue
}
