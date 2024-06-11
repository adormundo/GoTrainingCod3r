package main

import (
	"log"
	"net/http"
)

func main() {
	// Iniciando o servidor de arquivos estáticos
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Iniciando o servidor na porta 3000
	port := ":3000"
	log.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
