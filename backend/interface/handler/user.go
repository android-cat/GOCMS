package handler

import (
	"go-cms/domain/model"
	"go-cms/usecase"
	"log"
	"net/http"
	"strconv"

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

func (uh *UserHandler) VerifyEmail(ctx *gin.Context) {
    token := ctx.Query("token")
    if token == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Verification token is required"})
        return
    }

    err := uh.UserUseCase.VerifyEmail(token)
    if err != nil {
        log.Println(err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "メール認証に成功しました。ログインページに移動してログインしてください。"})
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := uh.UserUseCase.UpdateUser(&user); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := uh.UserUseCase.DeleteUserByID(uint(id)); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
