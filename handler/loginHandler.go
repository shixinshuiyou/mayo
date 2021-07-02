package handler

import "github.com/gin-gonic/gin"

func UserLogin(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0,"login success")
}

func UserLogout(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0,"logout success")
}

func UserInfo(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0,"logout success")
}
