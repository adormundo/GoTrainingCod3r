package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usuarioService := NewUsuarioService(db)

	http.HandleFunc("/usuarios/", UsuarioHandle(usuarioService))

	log.Println("Servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
