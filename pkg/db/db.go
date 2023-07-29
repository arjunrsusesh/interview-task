package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	db *sql.DB
)

func DBConnection(user, dbname, password, host, port string) error {
	config := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		user,
		dbname,
		password,
		host,
		port,
	)

	log.Println(config)
	var err error
	db, err = sql.Open("mysql", config)
	if err != nil {
		return errors.Wrap(err, "could not open database")

	}

	err = db.Ping()
	if err != nil {
		return errors.Wrap(err, "failed to ping DB")
	}

	log.Println("Succesfully connected to the database")
	return err
}

func GetDBConnection() *sql.DB {
	return db
}

func CloseDBConnection() {
	db.Close()
}
