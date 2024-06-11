package main

import "fmt"

// notaParaConceito converte uma nota numérica para um conceito de letra.
func notaParaConceito(n float64) string {
	nota := int(n) // Converte a nota para um inteiro
	switch nota {
	case 10:
		fallthrough // ignora o case 10 e pula para o próximo
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	// Testando a função notaParaConceito com diferentes notas
	fmt.Println(notaParaConceito(10))   // Output: A
	fmt.Println(notaParaConceito(9))    // Output: A
	fmt.Println(notaParaConceito(8))    // Output: B
	fmt.Println(notaParaConceito(7))    // Output: B
	fmt.Println(notaParaConceito(6))    // Output: C
	fmt.Println(notaParaConceito(5))    // Output: C
	fmt.Println(notaParaConceito(4))    // Output: D
	fmt.Println(notaParaConceito(3))    // Output: D
	fmt.Println(notaParaConceito(2))    // Output: E
	fmt.Println(notaParaConceito(1))    // Output: E
	fmt.Println(notaParaConceito(0))    // Output: E
	fmt.Println(notaParaConceito(-1))   // Output: Nota inválida
	fmt.Println(notaParaConceito(10.5)) // Output: A
	fmt.Println(notaParaConceito(7.9))  // Output: B
}
