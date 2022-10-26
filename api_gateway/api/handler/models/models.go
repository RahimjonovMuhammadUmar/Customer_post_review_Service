package models

type UserRegister struct {
	FirstName string //`json:"first_name"`
	LastName  string //`json:"last_name"`
	Username  string //`json:"username"`
	Password  string //`json:"password"`
	Email     string //`json:"email"`
}

type Error struct {
	Error string //`json:"error"`
}
