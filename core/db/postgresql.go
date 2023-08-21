package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "library"
)

type Postgres struct {
	Connection *sql.DB
}

func NewDatabase() (*Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbConnection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := dbConnection.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{Connection: dbConnection}, nil
}

// Close closes the database connection
func (d *Postgres) Close() {
	d.Connection.Close()
}

func (d *Postgres) Setup() {
	err := d.CreateAuthorTable()
	if err != nil {
		log.Fatal(err)
	}

	err = d.CreateBookTable()
	if err != nil {
		log.Fatal(err)
	}

	err = d.CreateBookAuthorJunctionTable()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Postgres) CreateAuthorTable() error {
	createAuthorTableSQL :=
		`
      CREATE TABLE IF NOT EXISTS author (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    )`

	_, err := d.Connection.Exec(createAuthorTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *Postgres) CreateBookTable() error {
	createBookTableSQL :=
		`
      CREATE TABLE IF NOT EXISTS book (
        id SERIAL PRIMARY KEY,
        name VARCHAR(150) NOT NULL,
        edition INTEGER NOT NULL,
        publication_year INTEGER NOT NULL
      )
  `
	_, err := d.Connection.Exec(createBookTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *Postgres) CreateBookAuthorJunctionTable() error {
	createBookAuthorJunctionSQL :=
		`
      CREATE TABLE IF NOT EXISTS author_book (
        book_id INTEGER REFERENCES book(id),
        author_id INTEGER REFERENCES author(id),
        PRIMARY KEY (book_id, author_id)
      )
    `

	_, err := d.Connection.Exec(createBookAuthorJunctionSQL)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
