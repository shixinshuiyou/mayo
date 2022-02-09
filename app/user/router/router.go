package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/mayo/app/user/handler"
	"github.com/shixinshuiyou/mayo/config"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

func Register() *gin.Engine {
	srvName := config.SrvActionName
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(tracer.Jaeger(srvName))
	// r.Use(metric.GinMiddleWare(selfConfig.SrvCube))
	// r.Use(gin2micro.TracerWrapper)
	r.Use(CorsMiddleware())
	r.Use(CheckSessionMiddleware())
	rg := r.Group(strings.Split(srvName, ".")[3])
	initAction(rg)
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL, c.Request.Host, c.Params)
	})
	return r
}

func CheckSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func initAction(r *gin.RouterGroup) {
	actionGroup := r.Group("action")
	actionGroup.GET("log", handler.UserLogin)
	actionGroup.POST("register", handler.UserLogout)
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
