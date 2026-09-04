package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"msg": "ok"})
			})
		}
	}
}
