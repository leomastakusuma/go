package main

import (
	Controller "github.com/second/controller"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB 
func main() {
	r := mux.NewRouter()
	Controller.Routers(r)
	http.ListenAndServe(":9090", r)

}
