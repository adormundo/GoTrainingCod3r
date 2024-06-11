package main

import "fmt"

// Definição da interface esportivo
type esportivo interface {
	ligarTurbo()
}

// Definição da interface luxuoso
type luxuoso interface {
	fazerBaliza()
}

// Definição da interface esportivoLuxuoso, que incorpora as interfaces esportivo e luxuoso
type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

// Struct bmw7 que implementa as interfaces esportivo e luxuoso
type bmw7 struct{}

// Implementação do método ligarTurbo para a struct bmw7
func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo ligado, vruuuuuuuum")
}

// Implementação do método fazerBaliza para a struct bmw7
func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	// Criando uma variável do tipo esportivoLuxuoso e atribuindo uma instância de bmw7
	var b esportivoLuxuoso = bmw7{}

	// Chamando métodos através da interface esportivoLuxuoso
	b.ligarTurbo()
	b.fazerBaliza()
}
