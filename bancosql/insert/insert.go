package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// openDB abre uma conexão com o banco de dados MySQL.
func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir conexão com o banco de dados: %v", err)
	}
	return db, nil
}

// inserirUsuario insere um novo usuário no banco de dados.
func inserirUsuario(db *sql.DB, nome string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	if err != nil {
		return 0, fmt.Errorf("falha ao preparar statement SQL: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(nome)
	if err != nil {
		return 0, fmt.Errorf("falha ao executar query SQL: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("falha ao obter ID do usuário inserido: %v", err)
	}

	linhas, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("falha ao obter número de linhas afetadas: %v", err)
	}

	log.Printf("Usuário %s inserido com sucesso. ID: %d, Linhas afetadas: %d\n", nome, id, linhas)

	return id, nil
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inserir usuários
	idMaria, err := inserirUsuario(db, "Maria")
	if err != nil {
		log.Fatalf("falha ao inserir Maria: %v", err)
	}
	fmt.Println("ID da Maria:", idMaria)

	idJoao, err := inserirUsuario(db, "João")
	if err != nil {
		log.Fatalf("falha ao inserir João: %v", err)
	}
	fmt.Println("ID do João:", idJoao)

	idPedro, err := inserirUsuario(db, "Pedro")
	if err != nil {
		log.Fatalf("falha ao inserir Pedro: %v", err)
	}
	fmt.Println("ID do Pedro:", idPedro)
}
