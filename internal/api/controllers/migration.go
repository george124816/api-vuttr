package controllers

import (
	"log"
	"time"
	"database/sql"
	"github.com/george124816/my-own-migration/migrate"
)


func Migrate(db *sql.DB){
	var updated bool = false
	for !updated{
		log.Println("trying to migrate.")
		updated = migrate.Migrate(database, 2)
		time.Sleep(5 * time.Second)
	}
}