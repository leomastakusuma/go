package Controller

import (
	"encoding/json"
	"log"
	 "os"
	"net/http"
	"../library"
	"../model"
	"github.com/gorilla/mux"
)

func Routers(r *mux.Router) {
	r.HandleFunc("/user", Myprofile).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
}

func Myprofile(w http.ResponseWriter, r *http.Request) {
	var users Model.Users
	var arr_user []Model.Users

	db := library.GetDB()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())	
		} else {
			arr_user = append(arr_user, users)
		}
	}
	response := make(map[string] interface{})
	response = library.Message(true, "list data") 
	response["data"] = arr_user
	library.Respond(w, response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db := library.GetDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(first_name,last_name) VALUES(?,?)")
	if err != nil {
		println("Exec err:", err.Error())
	}

	var user Model.Users
	var arr_user []Model.Users
	errs 			:= json.NewDecoder(r.Body).Decode(&user)
	arr_user 		= append(arr_user, user)
	if errs != nil {
		println("Exec err:", err.Error())
	}

	response 	:= make(map[string] interface{})
	firstName 	:= user.FirstName
	lastName 	:= user.LastName
	_, err 	= stmt.Exec(firstName, lastName)

	if err != nil {
		println("Exec err:", err.Error())
		response = library.Message(false, "error create user") 
		library.Respond(w, response)
    } else {
        if err != nil {
            println("Error:", err.Error())
		} 		
		response = library.SuccessInsert() 
		library.Respond(w, response)
    }
}

