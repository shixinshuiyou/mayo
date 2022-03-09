package srv

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

type SnowID struct {
	TraceContext context.Context // 需要携带TraceID
}

func NewIDSrv(ctx *gin.Context) (sid *SnowID) {
	sid = new(SnowID)
	if c, ok := tracer.ContextWithSpan(ctx); ok {
		sid.TraceContext = c
	} else {
		sid.TraceContext = ctx
	}
	return
}

func (sid *SnowID) GetSnowflakeID() (int64, error) {
	if iDService == nil {
		log.Logger.Warnf("未初始化proto协议client!!!")
		InitProtoService(client.DefaultClient)
	}

	resp, err := iDService.GetSnowflakeID(sid.TraceContext, &proto.SnowIDReq{})
	if err != nil {
		log.Logger.Errorf("get snowflakeID err:%s", err)
		return 0, nil
	}
	return resp.ID, nil
}
