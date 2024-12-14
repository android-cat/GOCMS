package router

import (
	"go-cms/interface/handler"
	"go-cms/interface/handler/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// 静的ファイルの提供
	router.Static("/static", "./ui/html")

	// セッションミドルウェアを適用
	router.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))

	// ハンドラーのインスタンスを作成
	loginHandler := handler.NewLoginHandler(db)
	registerHandler := handler.NewRegisterHandler(db)
	logoutHandler := handler.NewLogoutHandler()
	userHandler := handler.NewUserHandler(db)

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("./ui/html/top.html")
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.File("./ui/html/auth/login.html")
	})
	router.POST("/login", loginHandler.Login)

	router.GET("/register", func(ctx *gin.Context) {
		ctx.File("./ui/html/auth/register.html")
	})
	router.POST("/register", registerHandler.Register)

	router.POST("/logout", logoutHandler.Logout)
	
	authGroup := router.Group("/admin")
	authGroup.Use(middleware.IsLoggedIn())
	{
		authGroup.GET("/dashboard", func(ctx *gin.Context) {
			ctx.File("./ui/html/admin.html")
		})
		authGroup.GET("/userinfo", userHandler.GetUserInfo)
		authGroup.GET("/users", userHandler.GetAllUsers)
	}

	return router
}
