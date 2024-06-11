package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrindo conexão com o banco de dados
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Iniciando uma transação
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Preparando e executando as operações de inserção
	stmt, err := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(2000, "Bia")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = stmt.Exec(2001, "Carlos")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = stmt.Exec(1, "Tiago") // tentativa de inserir chave duplicada
	if err != nil {
		tx.Rollback()
		log.Printf("Erro ao executar a inserção: %v\n", err)
	}

	// Commit da transação
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserções realizadas com sucesso!")
}
