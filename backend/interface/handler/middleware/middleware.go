package middleware

import (
	"go-cms/domain/model"

	"os"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// isAuthenticatedは認証状態を確認する関数
func isAuthenticated(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	tokenString := session.Get("token")
	if tokenString == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return false
	}

	claims := &model.Claims{}
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.ParseWithClaims(tokenString.(string), claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return false
	}
	return true
}

func IsLoggedIn() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // ログインチェックのロジックを記述
        if !isAuthenticated(ctx) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}