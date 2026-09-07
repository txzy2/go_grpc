package auth

import "github.com/gin-gonic/gin"

func SetupRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "ok"})
		})

		auth.POST("/reg", func(c *gin.Context) {
			// TODO: Сделать сервис под создание пользователя
			c.JSON(200, gin.H{"msg": "ok"})
		})
	}
}
