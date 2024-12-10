package middleware

import (
	"github.com/gin-gonic/gin"
)

// isAuthenticatedは認証状態を確認する関数
func isAuthenticated(ctx *gin.Context) bool {
	// 認証チェックのロジック（例: セッションやトークンの確認）
	// 本例では固定値で「認証失敗」を返す
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return false
	}
	return true
}

func IsLoggedIn() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // ログインチェックのロジックを記述
        if !isAuthenticated(ctx) {
            return
        }
        ctx.Next()
    }
}
