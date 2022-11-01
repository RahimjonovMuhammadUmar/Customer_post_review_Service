package models

type CustomerRegister struct {
	FirstName string //`json:"first_name"`
	LastName  string //`json:"last_name"`
	Bio       string //`json:"username"`
	Password  string //`json:"password"`
	Email     string //`json:"email"`
	Addresses []AddressRequest
}

type Error struct {
	Error string //`json:"error"`
}

type CustomerDataToSave struct {
	FirstName string //`json:"first_name"`
	LastName  string //`json:"last_name"`
	Bio       string //`json:"username"`
	Password  string //`json:"password"`
	Email     string //`json:"email"`
	Code      int
	Addresses []AddressRequest
}

type CustomerResponse struct {
	FirstName string //`json:"first_name"`
	LastName  string //`json:"last_name"`
	Bio       string //`json:"username"`
	Password  string //`json:"password"`
	Email     string //`json:"email"`
	Addresses []AddressResponse
}

type AddressRequest struct {
	Street       string
	House_number int32
}

type AddressResponse struct {
	Id           int32
	Street       string
	House_number int32
}
type ResponseError struct {
	Error interface{} `json:"error"`
}

type ServerError struct {
	Status string `json:"status"`
	Message string `json:"message"`
}