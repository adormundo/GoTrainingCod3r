package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// dbConfig contém a configuração do banco de dados.
type dbConfig struct {
	user     string
	password string
	host     string
	port     string
	dbName   string
}

// newDBConfig cria uma nova configuração de banco de dados com base em variáveis de ambiente.
func newDBConfig() dbConfig {
	return dbConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
	}
}

// connectDB conecta-se ao banco de dados MySQL.
func connectDB(cfg dbConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.user, cfg.password, cfg.host, cfg.port, cfg.dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
	}
	return db, nil
}

// execQuery executa uma query SQL no banco de dados e retorna um erro, se houver.
func execQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao executar query: %v", err)
	}
	return nil
}

// setupDatabase cria e configura o banco de dados.
func setupDatabase(db *sql.DB) error {
	// Criar o banco de dados 'cursogo' se não existir.
	if err := execQuery(db, "CREATE DATABASE IF NOT EXISTS cursogo"); err != nil {
		return err
	}

	// Usar o banco de dados 'cursogo'.
	if err := execQuery(db, "USE cursogo"); err != nil {
		return err
	}

	// Criar a tabela 'usuarios' se não existir.
	const createTableUser = `CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`
	if err := execQuery(db, createTableUser); err != nil {
		return err
	}

	log.Println("Banco de dados configurado com sucesso!")
	return nil
}

func main() {
	cfg := newDBConfig()

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	if err := setupDatabase(db); err != nil {
		log.Fatalf("erro ao configurar o banco de dados: %v", err)
	}

	log.Println("Aplicação finalizada.")
}
