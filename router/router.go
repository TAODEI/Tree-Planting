package router

import (
	"TreePlanting/handler/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	r.Use(Cors())
	r.POST("/api/user", user.Login)
	r.POST("/api/content", user.PushContent)
	r.GET("/api/content", user.GetContent)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				// core.Logger.Error("Panic info is: %v", err)
				// core.Logger.Error("Panic info is: %s", debug.Stack())
			}
		}()

		c.Next()
	}
}

//
// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		origin := c.Request.Header.Get("Origin") // 请求头部
// 		if origin != "" {
// 			// 可将将* 替换为指定的域名
// 			c.Header("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 		}
//
// 		if method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)
// 		}
//
// 		c.Next()
// 	}
// }
