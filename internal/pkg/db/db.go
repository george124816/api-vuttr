package db

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "mysql"
	port = 3306
	user = "root"
	password = "secret"
	dbname = "main"
)

func ConnectDB() *sql.DB {
	
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Printf("Erro", err)
	}

	return db
}