package main

import "fmt"

/*
   Em Go, os operadores ++ e -- só podem ser usados de forma pós-fixada
   (ou seja, x++ e y-- são válidos, mas ++x e --y não são). Isso é diferente de
   outras linguagens como C ou C++, onde os operadores de pré-incremento e
   pré-decremento são permitidos.
*/

func main() {
	x := 1
	y := 2

	// apenas posfix
	x++ // x += 1 || x = x + 1
	fmt.Println(x)

	y-- // y -= 1 || y = y - 1
	fmt.Println(y)

}
