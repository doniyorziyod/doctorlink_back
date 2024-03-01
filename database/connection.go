package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)


func Connect() (*sqlx.DB, error) {
    db, err := sqlx.Connect("postgres", "user=doniyorziyod dbname=doctorlink sslmode=disable password=7355950d host=localhost")

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }

    return db, err
}
