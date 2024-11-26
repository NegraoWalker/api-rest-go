package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Configuração da conexão com o Postgres
	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=api sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// Criação da tabela "users"
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Não foi possível criar a tabela 'users': %v", err)
	}

	// Criação da tabela "events"
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		user_id INTEGER REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Não foi possível criar a tabela 'events': %v", err)
	}
}
