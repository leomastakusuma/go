package Controller

import (
	"encoding/json"
	"log"
	"net/http"
	"../library"
	"../model"
	"github.com/gorilla/mux"
)

type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}

func Routers(r *mux.Router) {
	r.HandleFunc("/user", Myprofile).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/test", testValidate).Methods("POST")
}

func Myprofile(w http.ResponseWriter, r *http.Request) {
	var users Model.Users
	var arr_user []Model.Users

	db := library.GetDB()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from person order by id desc")
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
	response = library.Message("list data")
	response["data"] = []int{}
	if(len(arr_user) > 0){
		response["data"] = arr_user
	}
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
	errs 	:= json.NewDecoder(r.Body).Decode(&user)
	if errs != nil {
		println("Exec err:", err.Error())
	}
	arr_user 		= append(arr_user, user)

	response 	:= make(map[string] interface{})
	firstName 	:= user.FirstName
	lastName 	:= user.LastName
	resp, err 	:= stmt.Exec(firstName, lastName)
	
	ID, err := resp.LastInsertId()
	bytes, err := json.Marshal(Model.Users{
		Id:ID,
        FirstName: user.FirstName,
        LastName: user.LastName,
	})
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
        panic(err)
    }

	if err != nil {
		println("Exec err:", err.Error())
		response = library.Message( "error create user") 
		library.Respond(w, response)
    } else {
        if err != nil {
            println("Error:", err.Error())
		} 		
		response = library.SuccessInsert() 
		response["data"]= data
		library.Respond(w, response)
    }
}

func testValidate(w http.ResponseWriter, r *http.Request){
	// rules := govalidator.MapData{
	// 	"firstname": []string{"required"},
	// 	"lastname" : []string{"required", "min:4", "max:20", "email"},
	// }

	// opts := govalidator.Options{
	// 	Request:         r,        // request object
	// 	Rules:           rules,    // rules map
	// 	RequiredDefault: true,     // all the field to be pass the rules
	// }
	// v := govalidator.New(opts)
	// e := v.Validate()
	// err := map[string]interface{}{"validationError": e}
	// w.Header().Set("Content-type", "applciation/json")
	// json.NewEncoder(w).Encode(err)
	// kucing := map[string]interface{} []
	// test := map[string]interface{} {"status" : "e", "message" : "s","data" : ""}
	// json.NewEncoder(w).Encode(kucing)

	var user Model.Repositories
	b, _ := json.Marshal(user)
    var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err != nil {
        panic(err)
    }
	json.NewEncoder(w).Encode(dat)

}