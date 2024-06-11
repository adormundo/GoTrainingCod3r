package main

import (
	"fmt"
	"math/rand"
	"time"
)

// numeroAleatorio retorna um número aleatório entre 0 e 9.
func numeroAleatorio() int {
	// Cria uma nova fonte de aleatoriedade baseada no timestamp atual.
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	// Retorna um número aleatório entre 0 e 9.
	return r.Intn(10)
}

func main() {
	/*
		O uso de blocos de inicialização em if-else é útil para limitar o escopo
		de variáveis temporárias e é uma prática recomendada quando essas
		variáveis não são necessárias fora do bloco.
	*/

	// Bloco de inicialização dentro de um if-else
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!", i)
	} else {
		fmt.Println("Perdeu!", i)
	}
}
