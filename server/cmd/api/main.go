package main

import (
	"fmt"
	"kyimmQ/student_api/internal/repository"
	"kyimmQ/student_api/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const port = 8080

type application struct {
	DBInfo string
	DB     repository.DatabaseRepo
}

func main() {

	// Get MySQL connection information from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// main app
	var app application
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	app.DBInfo = connectionString
	// create a database object which can be used
	// to connect with database.
	conn, err := app.connectToDB()
	if err != nil {
		panic(err)
	}
	app.DB = &dbrepo.MySQLDBRepo{DB: conn}
	defer app.DB.Connect().Close()

	log.Println("Starting application on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
