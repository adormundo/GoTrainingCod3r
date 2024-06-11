package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// Usando a interface vazia (interface{})
	var coisa interface{}
	fmt.Println(coisa) // <nil> (zero value de interface{} é nil)

	coisa = 3
	fmt.Println(coisa) // 3

	// Usando uma interface definida (type dinamico interface{})
	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2) // Opa

	coisa2 = true
	fmt.Println(coisa2) // true

	// Agora coisa2 pode armazenar qualquer tipo que implemente dinamico
	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2) // {Golang: Explorando a Linguagem do Google}
}
