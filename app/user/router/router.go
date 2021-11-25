package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/mayo-dev/app/user/handler"
)

func Register() *gin.Engine {
	r := gin.Default()
	// r.Use(metric.GinMiddleWare(selfConfig.SrvCube))
	// r.Use(gin2micro.TracerWrapper)
	r.Use(CorsMiddleware())
	r.Use(CheckSessionMiddleware())
	rg := r.Group("action")
	initCube(rg)
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL, c.Request.Host, c.Params)
	})
	return r
}

func CheckSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func initCube(r *gin.RouterGroup) {
	r.GET("log", handler.UserLogin)
	r.GET("test", handler.KuaiApitest)
}

// 跨域问题
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		var isAccess = true
		if isAccess {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Session-Id, Session-Scene, Qh-Mid, Version-Info")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}

		c.Next()
	}
}
