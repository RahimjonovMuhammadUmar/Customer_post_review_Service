package utils

import "net/mail"


func IsValidMail(email string) (string,error) {
	add,err := mail.ParseAddress(email)
	if err != nil {
		return "",err
	}

	return add.Address,nil

}