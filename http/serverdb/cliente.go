package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type UsuarioService struct {
	db *sql.DB
}

func NewUsuarioService(db *sql.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

func (us *UsuarioService) ObterPorID(id int) (*Usuario, error) {
	var u Usuario
	err := us.db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Nome)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (us *UsuarioService) ObterTodos() ([]Usuario, error) {
	rows, err := us.db.Query("SELECT id, nome FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		err := rows.Scan(&usuario.ID, &usuario.Nome)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (us *UsuarioService) Fechar() {
	us.db.Close()
}

func UsuarioHandle(us *UsuarioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := strings.TrimPrefix(r.URL.Path, "/usuarios/")
		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ID inválido")
			return
		}

		switch {
		case r.Method == http.MethodGet && id > 0:
			usuarioPorID(w, us, id)
		case r.Method == http.MethodGet:
			UsuarioTodos(w, us)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Desculpa... :(")
		}
	}
}

func usuarioPorID(w http.ResponseWriter, us *UsuarioService, id int) {
	u, err := us.ObterPorID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Usuário não encontrado")
		return
	}

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao serializar usuário")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func UsuarioTodos(w http.ResponseWriter, us *UsuarioService) {
	usuarios, err := us.ObterTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao obter usuários")
		return
	}

	jsonBytes, err := json.Marshal(usuarios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao serializar usuários")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
