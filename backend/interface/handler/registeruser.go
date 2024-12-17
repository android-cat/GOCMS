package handler

import (
	"log"
	"net/http"
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go-cms/usecase"
	"go-cms/infrastructure/email"
)

type RegisterHandler struct {
	RegisterUserUseCase *usecase.RegisterUserUseCase
}

func NewRegisterHandler(db *gorm.DB) *RegisterHandler {
	return &RegisterHandler{
		RegisterUserUseCase: usecase.NewRegisterUserUseCase(db),
	}
}

func generateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {
	userEmail := ctx.PostForm("email")
	password := ctx.PostForm("password")
	name := ctx.PostForm("name")

	if userEmail == "" || password == "" || name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email, password, and name are required"})
		return
	}

	token, err := generateToken()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate verification token"})
		return
	}

	err = email.SendVerificationEmail(userEmail, token)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	err = rh.RegisterUserUseCase.Register(userEmail, password, name, token)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully. Please check your email to verify your account."})
}
