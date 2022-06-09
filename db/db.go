package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaAoBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=climaRecife password=info@1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
