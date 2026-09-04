package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.GET("", func(c *gin.Context) {
					c.JSON(200, gin.H{"msg": "ok"})
				})
			}
		}

	}
}
