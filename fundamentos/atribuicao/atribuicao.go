package main

import "fmt"

// Função principal onde o programa inicia
func main() {

	// Declaração de uma variável do tipo byte
	var b byte = 3
	fmt.Println("Valor de b:", b)

	// Declaração de uma variável com inferência de tipo
	i := 3

	// Realizando operações aritméticas com i
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	fmt.Println("Valor de i após operações aritméticas:", i)

	// Declaração de múltiplas variáveis
	x, y := 1, 2
	fmt.Println("Valores iniciais de x e y:", x, y)

	// Troca dos valores de x e y
	x, y = y, x
	fmt.Println("Valores de x e y após a troca:", x, y)

	// Declaração de múltiplas variáveis com tipos explícitos
	var c, d int = 3, 4
	fmt.Println("Valores de c e d:", c, d)
}
