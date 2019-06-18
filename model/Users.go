
package Model

type Users struct {
	Id        int64 `form:"id" json:"id,omitempty"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

type Repositories struct {
	Data []Users 
}