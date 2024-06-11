package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir conexão com o banco de dados: %v", err)
	}
	return db, nil
}

func consultarUsuarios(db *sql.DB, id int) ([]usuario, error) {
	rows, err := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", id)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar consulta SQL: %v", err)
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var u usuario
		err := rows.Scan(&u.id, &u.nome)
		if err != nil {
			return nil, fmt.Errorf("falha ao ler resultado da consulta: %v", err)
		}
		usuarios = append(usuarios, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar nas linhas do resultado: %v", err)
	}

	return usuarios, nil
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	usuarios, err := consultarUsuarios(db, 5)
	if err != nil {
		log.Fatalf("falha ao consultar usuários: %v", err)
	}

	for _, u := range usuarios {
		fmt.Println(u)
	}
}
