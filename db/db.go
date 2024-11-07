package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Configuração da conexão com o Postgres
	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=api sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic("Não foi possível conectar ao banco de dados.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime TIMESTAMP NOT NULL,
    user_id INTEGER
	)
	`
	_, error := DB.Exec(createEventsTable)

	if error != nil {
		panic("Could not create events table: " + error.Error())
	}

}
