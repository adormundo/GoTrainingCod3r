package main

import "runtime/debug"

// f3 imprime a stack trace usando debug.PrintStack()
func f3() {
	debug.PrintStack()
}

// f2 chama f3
func f2() {
	f3()
}

// f1 chama f2
func f1() {
	f2()
}

func main() {
	// main chama f1
	f1()
}
