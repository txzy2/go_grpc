package v1

import (
	"second/handler/http/v1/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api/v1")
	{
		auth.SetupRoutes(api)
	}
}
