package main

import "fmt"

// incrementByValue recebe um valor int e incrementa esse valor localmente
func incrementByValue(n int) {
	n++
}

// incrementByPointer recebe um ponteiro para int e incrementa o valor apontado pelo ponteiro
func incrementByPointer(n *int) {
	*n++ // incrementa o valor apontado pelo ponteiro n
}

func main() {
	n := 1 // variável n do tipo inteiro

	// Por valor: passa uma cópia de n para incrementByValue
	incrementByValue(n)
	fmt.Println("Após incrementByValue:", n) // imprime 1 (não foi alterado dentro de incrementByValue)

	// Por referência: passa o endereço de memória de n para incrementByPointer
	incrementByPointer(&n)
	fmt.Println("Após incrementByPointer:", n) // imprime 2 (n foi incrementado dentro de incrementByPointer)
}
