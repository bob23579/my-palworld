package routers

import (
	"github.com/gin-gonic/gin"
	"my-palworld/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件添加 CORS 头
	r.Use(func(c *gin.Context) {
		// 允许所有来源的请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许特定的 HTTP 方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许特定的 HTTP 头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		// 允许浏览器发送 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是 OPTIONS 请求，则直接返回 200 状态码
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理请求
		c.Next()
	})
	// 用户相关接口
	// 登陆 退出  修改密码等操作
	authGroup := r.Group("/api/users")
	{
		authGroup.POST("/login", handlers.LoginHandler)
		authGroup.GET("/info", handlers.AuthMiddleware(), handlers.GetUserInfo)
		authGroup.GET("/modifyPassword", handlers.AuthMiddleware(), handlers.ModifyPassword)
	}
	gameGroup := r.Group("/api/game")
	{
		gameGroup.GET("/getGameStatus", handlers.AuthMiddleware(), handlers.GetGameStatusHandler)
		gameGroup.GET("/startGame", handlers.AuthMiddleware(), handlers.StartGameHandler)
		gameGroup.GET("/stopGame", handlers.AuthMiddleware(), handlers.StopGameHandler)
		gameGroup.GET("/restartGame", handlers.AuthMiddleware(), handlers.RestartGameHandler)
		gameGroup.GET("/updateGame", handlers.AuthMiddleware(), handlers.UpdateGameHandler)
		gameGroup.GET("/getGameConfig", handlers.AuthMiddleware(), handlers.GetGameConfigHandler)
	}

	// 服务器管理相关接口

	// 配置管理相关接口

	return r
}
