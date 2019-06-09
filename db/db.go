package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DB struct {
	Connection *sql.DB
	logger     *log.Logger
}

func Init() (*DB, error) {
	var db DB

	logFile, err := os.OpenFile("db.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return &DB{}, err
	}
	db.logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	db.Connection, err = sql.Open("sqlite3", "departments.db")
	if err != nil {
		return &DB{}, err
	}

	err = db.Connection.Ping()
	if err != nil {
		return &DB{}, err
	}

	err = db.firstCheckDatabase()
	if err != nil {
		return &DB{}, err
	}
	return &db, err
}

func (db *DB) firstCheckDatabase() error {
	var scripts = []string{createTableDepartment, createTableEmployee, createTablePosition}
	for i := range scripts {
		_, err := db.Connection.Exec(scripts[i])
		if err != nil {
			return err
		}
	}
	return nil
}
