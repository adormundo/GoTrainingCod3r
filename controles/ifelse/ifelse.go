package main

import "fmt"

// imprimirResultado avalia a nota e imprime se o aluno foi aprovado ou reprovado.
func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	// Testando a função imprimirResultado com diferentes notas
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
