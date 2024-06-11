package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsn                = "root:123456@/cursogo"
	createTableUsuario = `CREATE TABLE IF NOT EXISTS usuarios (
							id INT AUTO_INCREMENT PRIMARY KEY,
							nome VARCHAR(80)
						   )`
	insertUsuario = "INSERT INTO usuarios(id, nome) VALUES(?, ?)"
)

func main() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Erro ao iniciar transação:", err)
	}

	err = inserirUsuarios(tx)
	if err != nil {
		tx.Rollback()
		log.Fatal("Erro ao inserir usuários:", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Erro ao cometer transação:", err)
	}

	fmt.Println("Inserções realizadas com sucesso!")
}

func inserirUsuarios(tx *sql.Tx) error {
	stmt, err := tx.Prepare(insertUsuario)
	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(2000, "Bia")
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário Bia: %v", err)
	}

	_, err = stmt.Exec(2001, "Carlos")
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário Carlos: %v", err)
	}

	_, err = stmt.Exec(1, "Tiago") // tentativa de inserir chave duplicada
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário Tiago: %v", err)
	}

	return nil
}
