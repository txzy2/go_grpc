package v1

import (
	"second/handler/http/v1/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		auth.SetupRoutes(api)
	}
}
