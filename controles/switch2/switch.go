package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtém o tempo atual
	t := time.Now()

	// switch com expressão true omitida
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
