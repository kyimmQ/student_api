package main

import (
	"database/sql"
	"log"
)

func openDB(DBInfo string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DBInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DBInfo)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MySQL!")
	return connection, nil
}
