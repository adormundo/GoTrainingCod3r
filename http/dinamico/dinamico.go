package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	port = ":3000"
)

func horaCertaHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprint(w, "<h1>Hora certa: ", currentTime, "</h1>")
}

func setupRoutes() {
	http.HandleFunc("/", horaCertaHandler)
}

func main() {
	setupRoutes()

	server := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Servidor rodando em http://localhost%s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
