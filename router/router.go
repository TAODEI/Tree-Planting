package router

import (
	"TreePlanting/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	r.POST("/api/v1/user", user.Login)
	r.POST("/api/v1/content", user.PushContent)
}
