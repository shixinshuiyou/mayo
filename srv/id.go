package srv

import (
	"context"

	"github.com/gin-gonic/gin"
	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/tool/log"
)

type SnowID struct {
	TraceContext context.Context // 需要携带TraceID
}

func NewIDSrv(ctx *gin.Context) (sid *SnowID) {
	sid = new(SnowID)
	sid.TraceContext = ctx
	return
}

func (sid *SnowID) GetSnowflakeID() (int64, error) {
	resp, err := GetIDService(sid.TraceContext).GetSnowflakeID(sid.TraceContext, &proto.SnowIDReq{})
	if err != nil {
		log.Logger.Errorf("get snowflakeID err:%s", err)
	}
	return resp.ID, nil
}
