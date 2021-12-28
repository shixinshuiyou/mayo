package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailWithMsg(c *gin.Context, msg interface{}) {
	ResponseWithCode(c, CodeUnknownError, msg, nil)
}

func ResponseWithCode(c *gin.Context, msgCode int, msg interface{}, data interface{}) {
	if msg == nil {
		if val, ok := MsgCodeMap[msgCode]; ok {
			msg = val
		} else {
			msg = MsgCodeMap[-1]
		}
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    msgCode,
		"message": msg,
		"data":    data,
	})
}
