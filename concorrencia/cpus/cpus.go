package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Número de CPUs disponíveis
	fmt.Println("Número de CPUs disponíveis:", runtime.NumCPU())

	// Número de goroutines em execução
	fmt.Println("Número de goroutines em execução:", runtime.NumGoroutine())

	// Definir o número máximo de CPUs que podem ser usadas simultaneamente
	// runtime.GOMAXPROCS(2)

	// Estatísticas de alocação de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	/*
		Alloc se refere à quantidade total de bytes alocados e não liberados
		pela aplicação. Ou seja, representa a quantidade de memória que a aplicação
		está atualmente usando.
	*/
	fmt.Printf("Alloc = %v MiB(Mebibyte)\n", m.Alloc/1024/1024)

	// Versão do Go
	fmt.Println("Versão do Go:", runtime.Version())
}
