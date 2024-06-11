package main

import (
	"fmt"
	"math"
	"reflect"
	"unicode/utf8"
)

func main() {
	///// Números inteiros
	// Exibindo inteiros literais
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é do tipo", reflect.TypeOf(320000)) // int

	// Tipos inteiros sem sinal (só positivos): uint8(byte), uint16, uint32, uint64
	var b byte = 255  // byte é um apelido para o uint8
	var c uint8 = 255 // uint8 é um inteiro sem sinal de 8 bits
	fmt.Println("Tipo de b:", reflect.TypeOf(b), "Tipo de c:", reflect.TypeOf(c))

	// Tipos inteiros com sinal: int8, int16, int32, int64
	fmt.Println("Tamanho máximo int8:", math.MaxInt8)
	fmt.Println("Tamanho máximo int16:", math.MaxInt16)
	fmt.Println("Tamanho máximo int32:", math.MaxInt32)
	fmt.Println("Tamanho máximo int64:", math.MaxInt64)

	// Representação de caracteres Unicode (int32)
	var unicodetab rune = 'a' // rune é um alias para int32
	fmt.Println("O tipo rune é na verdade", reflect.TypeOf(unicodetab))

	///// Números reais (float32, float64)
	fmt.Println("Literal float é do tipo", reflect.TypeOf(32.0)) // float64 por padrão

	///// Boolean
	boleano := true
	fmt.Println("O tipo de boleano é", reflect.TypeOf(boleano))

	///// Strings
	str1 := "Olá meu nome é Álvaro"
	fmt.Println(str1 + "!")

	// len retorna a quantidade de bytes
	fmt.Println("O tamanho da String em bytes é", len(str1))

	// utf8.RuneCountInString retorna a quantidade de caracteres Unicode
	fmt.Println("O tamanho da String em caracteres é", utf8.RuneCountInString(str1))

	// Strings com múltiplas linhas
	str2 := `Olá
	meu nome
	é Álvaro`
	fmt.Println("A quantidade de bytes em str2 é", len(str2))
}
