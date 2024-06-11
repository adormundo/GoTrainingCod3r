package main

import "fmt"

// notaParaConceito converte uma nota numérica para um conceito de letra.
func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	// Testando a função notaParaConceito com diferentes notas
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(3.0))
	fmt.Println(notaParaConceito(2.0))
}
