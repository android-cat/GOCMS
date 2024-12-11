package repository

import "go-cms/domain/model"

type UserRepository interface {
	AllUsers() ([]model.User, error)
	GetUsersID(int) ([]model.User, error)
	GetUsersEmail(string) ([]model.User, error)
	CreateUser(*model.User) error
}

