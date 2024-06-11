package main

import (
	"fmt"
	"reflect"
)

// ImprimirAprovados imprime a lista de aprovados
func ImprimirAprovados(aprovados ...string) {
	fmt.Println("O tipo de aprovados é", reflect.TypeOf(aprovados))
	fmt.Println("Lista de Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovadosSlice := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	aprovadosArray := [...]string{"Julia", "Isa", "Alfredo", "Roberta"}

	ImprimirAprovados(aprovadosSlice...)
	ImprimirAprovados(aprovadosArray[:]...)
}
