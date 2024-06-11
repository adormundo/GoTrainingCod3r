package main

import "fmt"

// Carro representa um veículo genérico.
type Carro struct {
	Nome            string
	VelocidadeAtual int
}

// Ferrari representa um tipo específico de carro com turbo.
type Ferrari struct {
	Carro       // Campo anônimo para promover os campos de Carro
	TurboLigado bool
}

// LigarTurbo ativa o turbo da Ferrari.
func (f *Ferrari) LigarTurbo() {
	f.TurboLigado = true
}

func main() {
	f := Ferrari{}

	// Configurando a Ferrari diretamente
	f.Nome = "F40"
	f.VelocidadeAtual = 0
	f.LigarTurbo()

	fmt.Println(f)
	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f.Nome, f.TurboLigado)
}
