package main

import (
	"fmt"
)

// compras retorna se vai comprar a TV de 50", a TV de 32" e sorvete
// baseado nos resultados de dois trabalhos.
func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2 // ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n", tv50, tv32, sorvete, !sorvete)
}
