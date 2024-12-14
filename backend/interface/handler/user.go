package handler

import (
	"go-cms/usecase"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		UserUseCase: usecase.NewUserUseCase(db),
	}
}

func (uh *UserHandler) GetUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := uh.UserUseCase.GetUserByID(userID.(uint))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
    users, err := uh.UserUseCase.GetAllUsers()
    if err != nil {
		log.Println(err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
        return
    }

    ctx.JSON(http.StatusOK, users)
}