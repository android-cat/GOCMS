package usecase

import (
	"go-cms/domain/model"
	"go-cms/domain/repository"
	"go-cms/infrastructure/persistence"

	"gorm.io/gorm"
)

type RegisterUserUseCase struct {
    UserRepo repository.UserRepository
}

func NewRegisterUserUseCase(db *gorm.DB) *RegisterUserUseCase {
    return &RegisterUserUseCase{
        UserRepo: persistence.NewUserPersistence(db),
    }
}

func (uc *RegisterUserUseCase) Register(email, password, name string) error {
    // ユーザーを作成
    user := &model.User{
        Email:    email,
        Password: password,
        Name:     name,
    }

    // ユーザーをデータベースに保存
    return uc.UserRepo.CreateUser(user)
}