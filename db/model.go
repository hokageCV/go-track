package db

import (
	"database/sql"
	"os"

	"github.com/hokageCV/gotrack/utils"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

func InitializeDB() (*sql.DB, error) {
	// Check if the database file exists
	_, err := os.Stat("./db.sqlite3")
	if os.IsNotExist(err) {
		file, err := os.Create("./db.sqlite3")
		utils.CheckNilErr(err)
		file.Close()
	}

	// start connection with db
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	utils.CheckNilErr(err)

	// Ping the DB to check if the connection is successful
	err = db.Ping()
	utils.CheckNilErr(err)

	// create table if !exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Title TEXT,
			IsDone BOOLEAN DEFAULT FALSE
		)
	`)
	utils.CheckNilErr(err)

	return db, nil
}
