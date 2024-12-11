package handler

import (
	"go-cms/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterHandler struct {
	RegisterUserUseCase *usecase.RegisterUserUseCase
}

func NewRegisterHandler(db *gorm.DB) *RegisterHandler {
	return &RegisterHandler{
		RegisterUserUseCase: usecase.NewRegisterUserUseCase(db),
	}
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	name := ctx.PostForm("name")

	if email == "" || password == "" || name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email, password, and name are required"})
		return
	}

	err := rh.RegisterUserUseCase.Register(email, password, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
