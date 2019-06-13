package Controller

import (
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json"
		"../library"
		"../structs"
		"log"
		_ "github.com/go-sql-driver/mysql"


)

func Routers(r *mux.Router) {
    r.HandleFunc("/user/e", Myprofile).Methods("GET")
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



