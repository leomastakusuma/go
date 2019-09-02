package library

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "refit:refit2019@tcp(db.dev.antigravity.id:3306)/refit")
	if err != nil {
		panic(err)
	}
	return db
}
