package repository

import "BlacAi/internal/models"

type UserRepo interface {
	GetUserByEmail(email string)(models.UserDetails,error)
}


// func (r *UserRepo)GetUserByEmail(email string)(models.UserDetails,error){

// }