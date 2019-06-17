package library

import(
	"net/http"
	"encoding/json"
)
func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}


func SuccessInsert() (map[string]interface{}) {
	return map[string]interface{} {"status" : 200, "message" : "Success Insert"}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}