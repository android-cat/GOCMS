package usecase

import (
    "go-cms/domain/model"
    "go-cms/domain/repository"
    "go-cms/infrastructure/persistence"
    "gorm.io/gorm"
)

type UserUseCase struct {
    UserRepo repository.UserRepository
}

func NewUserUseCase(db *gorm.DB) *UserUseCase {
    return &UserUseCase{
        UserRepo: persistence.NewUserPersistence(db),
    }
}

func (uc *UserUseCase) GetUserByID(userID uint) (model.User, error) {
    return uc.UserRepo.GetUserByID(userID)
}

func (uc *UserUseCase) GetAllUsers() ([]model.User, error) {
    return uc.UserRepo.GetAllUsers()
}