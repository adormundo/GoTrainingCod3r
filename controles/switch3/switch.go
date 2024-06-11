package main

import "fmt"

// notaParaConceito converte uma nota numérica para um conceito de letra.
func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	// Testando a função notaParaConceito com diferentes notas
	fmt.Println(notaParaConceito(9.8)) // Output: A
	fmt.Println(notaParaConceito(8.9)) // Output: B
	fmt.Println(notaParaConceito(5.1)) // Output: C
	fmt.Println(notaParaConceito(3.0)) // Output: D
	fmt.Println(notaParaConceito(2.0)) // Output: E
}
