package usecase

import (
	"errors"
	"go-cms/domain/model"
	"go-cms/domain/repository"
	"go-cms/infrastructure/persistence"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUseCase struct {
	UserRepo repository.UserRepository
}

func NewLoginUseCase(db *gorm.DB) *LoginUseCase {
	return &LoginUseCase{
		UserRepo: persistence.NewUserPersistence(db),
	}
}

func (uc *LoginUseCase) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	users, err := uc.UserRepo.GetUsersEmail(email)
	if err != nil || len(users) == 0 {
		return "", errors.New("invalid email or password")
	}

	user := users[0]

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
