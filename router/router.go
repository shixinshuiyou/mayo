package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/mayo-dev/handler"
)

//Register 注册路由
func Register(port int) {
	r := gin.Default()
	// 心跳检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%d", port))
}

func InitUser(rg *gin.RouterGroup) {
	loginRouter := rg.Group("/user")
	loginRouter.POST("/login", handler.UserLogin)
	loginRouter.POST("/logout", handler.UserLogout)
	loginRouter.POST("/info", handler.UserInfo)

}

func InitMission(rg *gin.RouterGroup) {

}
