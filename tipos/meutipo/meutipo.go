package main

import "fmt"

// Definindo um novo tipo 'nota' que é um alias para float64.
// Isso permite que 'nota' seja tratada como um tipo distinto de float64.
type nota float64

// Método 'entre' associa ao tipo 'nota'.
// Este método verifica se o valor da nota está entre 'inicio' e 'fim'.
func (n nota) entre(inicio, fim float64) bool {
	// Convertendo 'n' para float64 e verificando se está dentro do intervalo [inicio, fim].
	return float64(n) >= inicio && float64(n) <= fim
}

// Função 'notaParaConceito' converte uma nota para um conceito.
// Ela recebe uma nota do tipo 'nota' e retorna o conceito correspondente.
func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	// Testando a função 'notaParaConceito' com diferentes notas.
	fmt.Println(notaParaConceito(9.8)) // Saída: A
	fmt.Println(notaParaConceito(6.9)) // Saída: C
	fmt.Println(notaParaConceito(2.1)) // Saída: E
}
