package web

import "github.com/gin-gonic/gin"

// Server gin web
func Server() {
	r := gin.Default()
	gin.DisableConsoleColor()
	r.GET("/user/:id", GetUserByID)
	r.GET("/sms/:mobile", GenSMSCode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
