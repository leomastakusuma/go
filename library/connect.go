package library

import (
	"database/sql"
)


var db *sql.DB
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
    return db
}