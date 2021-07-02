package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mayo-d/handler"
)

//Register 注册路由
func Register() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func initUser(rg *gin.RouterGroup) {
	loginRouter := rg.Group("/user")
	loginRouter.POST("/login", handler.UserLogin)
	loginRouter.POST("/logout",handler.UserLogout)
	loginRouter.POST("/info",handler.UserInfo)

}
