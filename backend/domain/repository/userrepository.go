package repository

import "go-cms/domain/model"

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(uint) (model.User, error)
	GetUserByEmail(string) (model.User, error)
	CreateUser(*model.User) error
	VerifyEmail(string) error
    UpdateUser(*model.User) error
    DeleteUserByID(uint) error
}

