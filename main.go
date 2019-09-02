package main

import (
	Controller "github.com/second/controller"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	Controller.Routers(r)
	log.Println("Starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", r))

}
