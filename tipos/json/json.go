package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Definição da struct produto
type produto struct {
	ID    int      `json:"id"` // Tags são usadas para definir o nome das chaves no JSON
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Convertendo struct para JSON
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))

	// Convertendo JSON para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(reflect.TypeOf(p2)) // Imprime o tipo da variável p2
}
