package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/log"
)

func UserLogin(ctx *gin.Context) {
	log.Logger.Debugf("======")
	ctx.AbortWithStatusJSON(0, "login success")
}

func UserLogout(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0, "logout success")
}

func UserInfo(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0, "logout success")
}
