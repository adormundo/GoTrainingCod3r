package main

import (
	"fmt"
	"reflect"
	"time"
)

// tipo retorna uma string que descreve o tipo da variável fornecida.
func tipo(any interface{}) string {
	switch any.(type) {
	case int, int8, int16, int32, int64:
		return "inteiro"
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return "inteiro sem sinal"
	case float32, float64:
		return "real"
	case complex64, complex128:
		return "complexo"
	case bool:
		return "booleano"
	case string:
		return "string"
	case [2]int, [3]int, [2]string, [3]string:
		return "array"
	case []int, []float64, []string:
		return "slice"
	case map[string]int, map[string]string:
		return "map"
	case chan int, chan string:
		return "canal"
	case *int, *float64, *string:
		return "ponteiro"
	case func():
		return "função"
	default:
		if reflect.TypeOf(any).Kind() == reflect.Struct {
			return "struct"
		}
		return "unknown"
	}
}

func main() {
	// Testando a função tipo com diferentes tipos de variáveis
	fmt.Println(tipo(2.3))        // Output: real
	fmt.Println(tipo(1))          // Output: inteiro
	fmt.Println(tipo("Opa"))      // Output: string
	fmt.Println(tipo(func() {}))  // Output: função
	fmt.Println(tipo(time.Now())) // Output: struct

	// Testando com outros tipos built-in
	fmt.Println(tipo(true))                                  // Output: booleano
	fmt.Println(tipo([]int{1, 2, 3}))                        // Output: slice
	fmt.Println(tipo([2]int{1, 2}))                          // Output: array
	fmt.Println(tipo(map[string]int{}))                      // Output: map
	fmt.Println(tipo(make(chan int)))                        // Output: canal
	fmt.Println(tipo(new(int)))                              // Output: ponteiro
	fmt.Println(tipo(interface{}(nil)))                      // Output: unknown
	fmt.Println(tipo(complex(1, 2)))                         // Output: complexo
	fmt.Println(tipo(byte(255)))                             // Output: inteiro sem sinal
	fmt.Println(tipo(uintptr(0)))                            // Output: inteiro sem sinal
	fmt.Println(tipo(struct{ Name string }{Name: "Gopher"})) // Output: struct
}
