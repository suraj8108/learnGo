package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Enter your Own Password"
	dbname   = "postgres"
)

func InitDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// dataSourceName := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		panic("Could Not connect to database")
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	DB.SetMaxOpenConns(10) // Number of connections allowed by an application to open
	DB.SetMaxIdleConns(5)  // Number of connections to keep alive if not being used.

	createTables(DB)
	return DB
}

func createTables(DB *sql.DB) {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS usersEvents (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
 		)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		fmt.Println(err)
		panic("Could Not create Event table")
	}

	creteEventTables := ` 
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY ,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES usersEvents (id)
	)
	`

	_, err = DB.Exec(creteEventTables)
	if err != nil {
		fmt.Println(err)
		panic("Could Not create Event table")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registration (
			id SERIAL PRIMARY KEY ,
			event_id INTEGER, 
			user_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events (id),
			FOREIGN KEY (user_id) REFERENCES usersEvents (id)
		)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		fmt.Println(err)
		panic("Could Not create Event table")
	}
}
