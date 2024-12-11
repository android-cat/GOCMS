package persistence

import (
	"go-cms/domain/model"
	"go-cms/domain/repository"

	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

func (up *userPersistence) AllUsers() ([]model.User, error) {
	users := []model.User{}
	res := up.db.Find(&users)
	if res.Error != nil {
		return []model.User{}, res.Error
	}
	return users, nil
}

func (up *userPersistence) GetUsersID(id int) ([]model.User, error) {
	users := []model.User{}
	up.db = up.db.Where("id = ?", id)
	res := up.db.Where("id = ?", id).Find(&users)
	if res.Error != nil {
		return []model.User{}, res.Error
	}
	return users, nil
}

func (up *userPersistence) GetUsersEmail(email string) ([]model.User, error) {
	users := []model.User{}
	up.db = up.db.Where("email = ?", email)
	res := up.db.Find(&users)
	if res.Error != nil {
		return []model.User{}, res.Error
	}
	return users, nil
}

func (up *userPersistence) CreateUser(user *model.User) error {
	if err := user.SetPassword(user.Password); err != nil {
		return err
	}

	res := up.db.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
