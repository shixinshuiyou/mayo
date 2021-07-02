package tool

import "github.com/gin-gonic/gin"

//GinResponse 所有Gin返回接口封装接口，要求所有Gin返回都必须使用此借口内的方法
type GinResponse interface {
	ResponseWithCode(ctx *gin.Context, msgCode int, msg interface{},data interface{})
}

func ResponseWithCode(ctx *gin.Context, msgCode int, msg interface{},data interface{}){

}