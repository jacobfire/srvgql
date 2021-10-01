package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Postgres *sql.DB

// Open connection and ping db
func (e *Environment) InitPostgres() (*sql.DB, error) {
	log.Println("Postgres Init!!!")
	log.Println(e.Databases["postgres"].Connection)

	db, err := sql.Open("postgres", e.Databases["postgres"].Connection)
	if err != nil {
		log.Println(err)

		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Postgres Ready!!!")

	Postgres = db
	return db, nil
}

// Close connection
func (e *Environment) PostgresClose() error {
	if err := Postgres.Close(); err != nil {
		return err
	}

	return nil
}

