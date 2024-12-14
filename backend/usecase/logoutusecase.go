package usecase

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

type LogoutUseCase struct{}

func NewLogoutUseCase() *LogoutUseCase {
    return &LogoutUseCase{}
}

func (uc *LogoutUseCase) Logout(ctx *gin.Context) {
    session := sessions.Default(ctx)
    session.Delete("token")
    session.Save()
    ctx.JSON(200, gin.H{"message": "Logout successful"})
}