package handler

import (
	"go-cms/usecase"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginHandler struct {
	LoginUseCase *usecase.LoginUseCase
}

func NewLoginHandler(db *gorm.DB) *LoginHandler {
	return &LoginHandler{
		LoginUseCase: usecase.NewLoginUseCase(db),
	}
}

func (lh *LoginHandler) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	tokenString, err := lh.LoginUseCase.Login(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	session := sessions.Default(ctx)
	session.Set("token", tokenString)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
