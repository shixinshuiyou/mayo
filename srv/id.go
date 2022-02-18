package srv

import (
	"context"

	"github.com/gin-gonic/gin"
	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

type SnowID struct {
	TraceContext context.Context // 需要携带TraceID
}

func NewIDSrv(ctx *gin.Context) (sid *SnowID) {
	sid = new(SnowID)
	sid.TraceContext = tracer.ContextWithTraceID(ctx)
	return
}

func (sid *SnowID) GetSnowflakeID() (int64, error) {
	resp, err := idService.GetSnowflakeID(sid.TraceContext, &proto.SnowIDReq{})
	if err != nil {
		log.Logger.Errorf("get snowflakeID err:%s", err)
	}
	return resp.ID, nil
}
