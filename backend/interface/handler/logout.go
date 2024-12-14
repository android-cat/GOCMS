package handler

import (
    "github.com/gin-gonic/gin"
    "go-cms/usecase"
)

type LogoutHandler struct {
    LogoutUseCase *usecase.LogoutUseCase
}

func NewLogoutHandler() *LogoutHandler {
    return &LogoutHandler{
        LogoutUseCase: usecase.NewLogoutUseCase(),
    }
}

func (lh *LogoutHandler) Logout(ctx *gin.Context) {
    lh.LogoutUseCase.Logout(ctx)
}