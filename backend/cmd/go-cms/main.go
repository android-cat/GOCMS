package main

import (
    "go-cms/interface/handler/router"
	"go-cms/infrastructure/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
    // データベース接続を取得
    db := database.NewDB()
    defer database.CloseDB(db)

    // ルーターを設定
    router := router.NewRouter(db)
    router.Run(":8080")

	// セッションストアを設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}