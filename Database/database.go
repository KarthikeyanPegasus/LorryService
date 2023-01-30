package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "user=tkzmgzyr password=lE5i7FCQS_WOQ2t9AVXu17JYhJ1xYqdy host=manny.db.elephantsql.com port=5432 dbname=tkzmgzyr sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	return db
}
