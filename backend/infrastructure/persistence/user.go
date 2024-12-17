package persistence

import (
	"go-cms/domain/model"
	"go-cms/domain/repository"
	"fmt"

	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

func (up *userPersistence) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	res := up.db.Session(&gorm.Session{}).Find(&users)
	if res.Error != nil {
		return []model.User{}, res.Error
	}
	return users, nil
}

func (up *userPersistence) GetUserByID(id uint) (model.User, error) {
	user := model.User{}
	res := up.db.Session(&gorm.Session{}).Where("id = ?", id).Find(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}
	return user, nil
}

func (up *userPersistence) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	res := up.db.Session(&gorm.Session{}).Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}
	return user, nil
}

func (up *userPersistence) CreateUser(user *model.User) error {
    return up.db.Transaction(func(tx *gorm.DB) error {
        // パスワードのハッシュ化
        if err := user.SetPassword(user.Password); err != nil {
            return fmt.Errorf("failed to hash password: %w", err)
        }
        
        // ユーザーの作成
        if err := tx.Create(user).Error; err != nil {
            return fmt.Errorf("failed to create user: %w", err)
        }
        
        return nil
    })
}

func (up *userPersistence) VerifyEmail(token string) error {
	user := model.User{}
	res := up.db.Session(&gorm.Session{}).Where("verification_token = ?", token).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	
	user.EmailVerified = true
	user.VerificationToken = ""
	
	if err := up.db.Session(&gorm.Session{}).Save(&user).Error; err != nil {
		return err
	}
	
	return nil
}
