package Controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../library"
	"../structs"
	"github.com/gorilla/mux"
)

func Routers(r *mux.Router) {
	r.HandleFunc("/user", Myprofile).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
}

func Myprofile(w http.ResponseWriter, r *http.Request) {
	var users structs.Users
	var arr_user []structs.Users
	var response structs.Response
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
	response.Status = true
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := library.GetDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO person(first_name,last_name) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	var user structs.Users
	var arr_user []structs.Users
	var response structs.Response

	errs := json.NewDecoder(r.Body).Decode(&user)
	arr_user = append(arr_user, user)
	if errs != nil {
		panic(err.Error())
	}

	firstName := user.FirstName
	lastName := user.LastName
	_, err = stmt.Exec(firstName, lastName)
	if err != nil {
		panic(err.Error())
	}
	response.Status = true
	response.Message = "Success"
	response.Data = arr_user

	json.NewEncoder(w).Encode(response)

}
