package service

import "golang.org/x/crypto/bcrypt"




func PasswordHashing(password string)(string,error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed),err
}