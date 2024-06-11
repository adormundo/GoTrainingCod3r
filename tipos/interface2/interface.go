package main

import "fmt"

// Definição da interface esportivo
type esportivo interface {
	ligarTurbo()
}

// Definição da struct ferrari
type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// Método ligarTurbo para a ferrari
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	// Exemplo com uma ferrari diretamente
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Exemplo com uma ferrari através da interface esportivo
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	// Imprimindo os carros
	fmt.Println(carro1) // Saída: {F40 true 0}
	fmt.Println(carro2) // Saída: &{F40 true 0}
}
